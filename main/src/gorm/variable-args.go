package main

import "fmt"

func main() {
	args("a", "b", "c")
}

func args(s ...string) (ss string) {

	testArr([]string{"1", "2", "3"}, s)
	return ss
}

func testArr(arr ...[]string) (res []string) {
	fmt.Println(arr)
	return
}
