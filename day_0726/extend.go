package main

/**
模拟面向对象类的继承
*/

type Flyable struct {
}

func (f Flyable) fly() {
	println("can fly")
}

type Walkable struct {
}

func (w Walkable) walk() {
	println("can walk")
}

type Bird struct {
	Flyable
	Walkable
}

type Human struct {
	Walkable
}

func main() {
	bird := Bird{}
	bird.fly()
	bird.walk()

	human := Human{}
	human.walk()
}
