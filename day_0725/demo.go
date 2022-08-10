package main

func main() {
	op := accumulator(1)
	println(op())
	println(op())
	op2 := accumulator(10)
	println(op2())
}

func accumulator(value int) func() int {
	return func() int {
		value++
		return value
	}
}
