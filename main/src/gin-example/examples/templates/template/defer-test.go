package main

import "fmt"

func main() {

	for i := 1; i < 10; i++ {
		defer c(i)
		fmt.Println("hhhh")
		if i == 1 {
			continue
		}
	}
}

func c(i int) {
	fmt.Println("===", i)
}
