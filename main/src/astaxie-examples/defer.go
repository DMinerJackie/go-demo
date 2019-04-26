package main

import "fmt"

func deferFunc() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

func main() {
	deferFunc()
}
