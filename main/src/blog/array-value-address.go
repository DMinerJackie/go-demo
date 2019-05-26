package main

import "fmt"

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	fmt.Printf("origin array address: %p \n", &arr)
	//passArray(arr)
	//fmt.Println(arr)
	passAddress(&arr)
	fmt.Println(arr)
}

func passArray(arr1 [5]int) {
	fmt.Printf("passed array address, arr1: %p \n", &arr1)
	fmt.Println(arr1)
	arr1[3] = 111
	fmt.Println("pass array arr1: ", arr1)
}

func passAddress(arr2 *[5]int) {
	fmt.Printf("passed array address, arr2: %p \n", arr2)
	fmt.Printf("passed array address, arr2: %p \n", &arr2)
	fmt.Println(arr2)
	arr2[3] = 111
	fmt.Println("pass array arr2: ", *arr2)
}
