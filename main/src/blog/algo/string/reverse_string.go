package main

import "fmt"

func main() {
	reverseString([]byte{'a', 'b', 'c', 'd', 'e'})
}

func reverseString(s []byte) {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	fmt.Printf("%s", s)
}
