package main

import (
	"fmt"
	"learning/entity"
)

func main() {
	//测试switch
	score(40)
	//测试交换
	fmt.Println(swap("hu", "chenhu"))
	//测试指针作为参数
	a := 1
	b := 2
	swapPtx(&a, &b)
	fmt.Println(a, b)
	//测试结构体
	authorInfo()
	//测试range
	testRange()
	//测试map
	testMap()
	//测试接口
	var animal Animal
	animal = new(Dog)
	animal.say()
	animal = new(Cat)
	animal.say()
	//测试channel
	nums := []int{1, 2, 3, 4, 5, 6}
	println(testChannel(nums))
}

/**
switch test
*/
func score(marks int) {
	var level = marks / 10
	var grade string
	switch level {
	case 9, 10:
		grade = "A"
	case 8:
		grade = "B"
	case 7, 6:
		grade = "C"
	default:
		grade = "D"
	}
	fmt.Printf("你的等级为%s\n", grade)
}

func swap(a string, b string) (string, string) {
	return b, a
}

func swapPtx(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func authorInfo() {
	var book entity.Book
	book.Title = "神印王座"
	book.BookId = 1
	book.Author = "唐家三少"
	book.Subject = "打怪升级"

	author := entity.Author{Name: "唐家三少", Age: 38, Sex: "男"}
	books := []entity.Book{book}
	author.Books = books
	fmt.Println(author)
}

func testRange() {
	str1 := "learning-go"
	str2 := "go语言学习"
	for _, s := range str1 {
		fmt.Printf("%c", s)
	}
	println()
	for _, s := range str2 {
		fmt.Printf("%c", s)
	}
	println()
}

func testMap() {
	numMap := make(map[string]string)
	numMap["one"] = "1"
	numMap["two"] = "2"
	for key, value := range numMap {
		fmt.Printf("%s:%s\n", key, value)
	}
}

type Animal interface {
	say()
}

type Dog struct {
}

func (dog Dog) say() {
	println("汪汪汪")
}

type Cat struct {
}

func (cat Cat) say() {
	println("喵喵喵")
}

func testChannel(nums []int) int {
	nums1 := nums[:len(nums)/2]
	nums2 := nums[len(nums)/2:]
	c := make(chan int)
	go sum(nums1, c)
	go sum(nums2, c)
	x, y := <-c, <-c
	return x + y
}

func sum(nums []int, c chan int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	c <- sum
}
