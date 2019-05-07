package main

import (
	"fmt"
	"strings"
)

type Foo struct {
	Name string
}

var bar = "hello Stark"

func returnValue() string {
	return bar
}

func returnPoint() *string {
	return &bar
}

func modifyNameByPoint(foo *Foo) {
	fmt.Printf("modifyNameByPoint foo: %v, address: %p \n", foo, &foo)
	foo.Name = "emmmm" + foo.Name
}

func nameToUpper(foo Foo) string {
	foo.Name = strings.ToUpper(foo.Name)
	fmt.Printf("nameToUpper foo: %v, address: %p \n", foo, &foo)
	return foo.Name
}

func (foo Foo) setName(name string) {
	fmt.Printf("setName foo: %v, address: %p\n", foo, &foo)
	foo.Name = name
}

func (foo *Foo) setNameByPoint(name string) {
	fmt.Printf("setNameByPoint foo: %v, address: %p\n", foo, &foo)
	foo.Name = name
}

func main() {
	//s1 := returnValue()
	//s2 := returnPoint()
	//
	//fmt.Println("s1: ", s1)
	//fmt.Println("s2: ", *s2)
	//
	//foo := Foo{Name : "jackie"}
	//fmt.Println("foo.name", foo.Name)
	//modifyNameByPoint(&foo)
	//fmt.Println("foo.name", foo.Name)
	//
	//fmt.Println("foo.name", nameToUpper(foo))
	//fmt.Println("foo.name", foo.Name)
	//
	//var bar = "hello Stark"
	//s11 := returnValue()
	//s22 := returnPoint()
	//
	////=======================
	//
	//fmt.Printf("bar: %v, address: %p \n", bar, &bar)
	//fmt.Printf("s11: %v, address: %p \n", s11, &s11)
	//fmt.Printf("s22: %v, address: %p \n", *s22, s22)

	////=======================
	//
	//foo := Foo{Name:"Stark"}
	//fmt.Printf("foo: %v, address: %p \n", foo, &foo)
	//
	//nameToUpper(foo)
	//fmt.Printf("foo: %v, address: %p \n", foo, &foo)

	////=======================
	//foo := &Foo{Name: "Stark"}
	//fmt.Printf("foo: %v, address: %p \n", foo, &foo)
	//
	//modifyNameByPoint(foo)
	//fmt.Printf("foo: %v, address: %p \n", foo, &foo)

	////=======================
	//foo := Foo{Name:"Stark"}
	//foo.setName("Jackie")
	//fmt.Println("foo.Name: ", foo.Name)
	//
	//foo.setNameByPoint("Jackie")
	//fmt.Println("foo.Name", foo.Name)

	//=======================
	foo := Foo{Name: "Stark"}
	fmt.Printf("src foo: %v, address: %p\n", foo, &foo)

	foo.setName("set name")
	fmt.Printf("by value foo: %v, address: %p\n", foo, &foo)

	foo.setNameByPoint("by point")
	fmt.Printf("by point foo: %v, address: %p\n", foo, &foo)

	var i = 1
	p := &i
	fmt.Printf("i: %v, p: %v\n", i, p)

	var f1 = &Foo{Name: "hello"}
	f2 := &f1
	fmt.Printf("f1: %v, f2: %v\n", f1, &f2)
}
