package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("d://")))
	_ = http.ListenAndServe(":8080", nil)
}

func test1() {
	channel := make(chan string)
	go producer("dog", channel)
	go producer("cat", channel)
	consumer(channel)
}

func producer(header string, channel chan<- string) {
	for {
		channel <- fmt.Sprintf("%s:%v", header, rand.Int31())
		time.Sleep(time.Second)
	}
}

func consumer(channel <-chan string) {
	for {
		str := <-channel
		fmt.Println(str)
	}
}
