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

	// int to string
	i := 100
	convertStr := strconv.Itoa(i)
	fmt.Println("convertStr: ", convertStr)

	// string to int
	convertInt, err := strconv.Atoi(convertStr)
	if err != nil {
		fmt.Println("string to int convert error")
	}
	fmt.Println("convertInt: ", convertInt)

	// string to int64
	convertInt64, err := strconv.ParseInt(convertStr, 10, 64)
	if err != nil {
		fmt.Println("convertInt64: ", convertInt64)
	}

	// int64 to string
	convertStr64 := strconv.FormatInt(convertInt64, 10)
	fmt.Println("convertStr64: ", convertStr64)

	convertFloat32, err := strconv.ParseFloat(convertStr, 32)
	if err != nil {
		fmt.Println("string to float32 error")
	}
	fmt.Println("convertFloat32: ", convertFloat32)

	convertFloat64, err := strconv.ParseFloat(convertStr, 64)
	if err != nil {
		fmt.Println("string to float64 error")
	}
	fmt.Println("convertFloat64; ", convertFloat64)

}
