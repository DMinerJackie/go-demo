package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) ReflectCallFunc() {
	fmt.Println("this is reflect call func")
}

func main() {
	user := User{1, "Jackie", 18}
	DoFieldAndMethod(user)
}

func DoFieldAndMethod(i interface{}) {
	getType := reflect.TypeOf(i)
	fmt.Println("type is: ", getType)

	getValue := reflect.ValueOf(i)
	fmt.Println("value is: ", getValue)

	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v \n", field.Name, field.Type, value)
	}

	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s, %v \n", m.Name, m.Type)
	}
}
