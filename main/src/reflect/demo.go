package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	t := reflect.TypeOf(x)
	fmt.Println("type: ", t)

	v := reflect.ValueOf(x)
	fmt.Println("value: ", v)

	var f float64 = 3.14
	v1 := reflect.ValueOf(f)
	var y float64 = v1.Interface().(float64)
	fmt.Println("y: ", y)

	var f1 float64 = 3.14
	v2 := reflect.ValueOf(&f1)
	v2.Elem().SetFloat(7.1)
	fmt.Println("v2: ", v2.Elem().Interface())
	fmt.Println("f1: ", f1)

	m := make(map[int64]int64)
	m[111] = 0
	m = nil
	m[22] = 22

}
