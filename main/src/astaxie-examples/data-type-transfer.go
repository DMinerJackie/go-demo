package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// string类型的slice转换为string字符串
	values := []string{"one", "two", "three", "four", "five"}
	fmt.Println("slice to string: ", strings.Join(values, "-"))

	// int类型的slice先转换为string类型的slice，在转换为string字符串
	valueInts := []int{10, 200, 3000}
	valuesText := []string{}

	for _, v := range valueInts {
		valuesText = append(valuesText, strconv.Itoa(v))
	}
	fmt.Println("valueText: ", valuesText)
	fmt.Println("valueText: ", strings.Join(valuesText, "-"))

}
