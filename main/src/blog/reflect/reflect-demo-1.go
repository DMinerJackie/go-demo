package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 1.2345

	fmt.Println(reflect.TypeOf(num))
	fmt.Println(reflect.ValueOf(num))
}
