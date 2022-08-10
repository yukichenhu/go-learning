package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Info struct {
	Username string
	Age      int
}

func main() {
	f, _ := os.Open("./info.json")
	defer f.Close()
	decoder := json.NewDecoder(f)
	info := Info{}
	_ = decoder.Decode(&info)
	fmt.Println(info)
}
