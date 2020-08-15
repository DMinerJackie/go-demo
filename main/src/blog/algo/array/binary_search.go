package main

import "fmt"

func main() {
	fmt.Println(4 >> 1)
}

func binarySearch(array []int, des int) int {
	low := 0
	high := len(array) - 1
	for low <= high {
		mid := low + (high-low)>>1
		if des == array[mid] {
			return mid
		} else if des < array[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
