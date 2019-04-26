package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"jackie", 18})
	fmt.Println(person{name: "hello", age: 10})
	fmt.Println(person{name: "world"})

	fmt.Println(&person{name: "go", age: 12})
	s := person{name: "wonderful", age: 15}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	ss := person{age: 11}
	fmt.Println(ss.name)
}
