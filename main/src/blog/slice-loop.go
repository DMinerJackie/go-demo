package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	for index, v := range s {
		fmt.Println(index, v)
	}

	var x, y []int
	for i := 0; i < 10; i++ {
		y = append(x, i)
		fmt.Printf("%d cap=%d\t %v\n", i, cap(y), y)
		x = y
	}
}
