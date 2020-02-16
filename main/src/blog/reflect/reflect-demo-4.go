package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 1.2345
	fmt.Println("old value is: ", num)

	pointer := reflect.ValueOf(&num)
	newValue := pointer.Elem()

	fmt.Println("type of pointer:", newValue.Type())
	fmt.Println("settability of pointer:", newValue.CanSet())

	newValue.SetFloat(2.3456)
	fmt.Println("new value of pointer:", num)

}
