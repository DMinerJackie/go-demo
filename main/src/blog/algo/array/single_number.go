package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(f())
}

func f() (video string) {
	video, err := "s", errors.New("hhh")
	if err != nil {

	}
	return
}

func singleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}
