package main

import "fmt"

type IGreeting interface {
	sayHello()
}

func sayHello(i IGreeting) {
	i.sayHello()
}

type Go struct {
}

func (g Go) sayHello() {
	fmt.Println("Hi, I am Go")
}

type PHP struct {
}

func (php PHP) sayHello() {
	fmt.Println("Hi, I am PHP")
}

func main() {
	golang := Go{}
	php := PHP{}

	sayHello(golang)
	sayHello(php)
}
