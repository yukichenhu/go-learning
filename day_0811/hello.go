package main

import (
	"flag"
	flag2 "learning/entity/flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "username", "your name")
}

func main() {
	flag.Parse()
	flag2.Hello(name)
}
