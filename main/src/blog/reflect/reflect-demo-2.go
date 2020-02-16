package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 1.2345

	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)

	conveterPointer := pointer.Interface().(*float64)
	conveterValue := value.Interface().(float64)

	fmt.Println(conveterPointer)
	fmt.Println(conveterValue)
}
