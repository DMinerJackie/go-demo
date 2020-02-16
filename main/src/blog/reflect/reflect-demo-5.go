package main

import (
	"fmt"
	"reflect"
)

type User1 struct {
	Id   int
	Name string
	Age  int
}

func (u User1) ReflectCallFuncHasArgs(name string, age int) {
	fmt.Println("ReflectCallFuncHasArgs name: ", name, ", age:", age, "and origal User.Name:", u.Name)
}

func (u User1) ReflectCallFuncNoArgs() {
	fmt.Println("ReflectCallFuncNoArgs")
}

func main() {
	user := User1{1, "Jackie", 20}

	getValue := reflect.ValueOf(user)
	methodValue := getValue.MethodByName("ReflectCallFuncHasArgs")
	args := []reflect.Value{reflect.ValueOf("Jackie"), reflect.ValueOf(20)}
	methodValue.Call(args)

	methodValue = getValue.MethodByName("ReflectCallFuncNoArgs")
	args = make([]reflect.Value, 0)
	methodValue.Call(args)
}
