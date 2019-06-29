package main

import "fmt"

func main() {
	fmt.Println("f1 result: ", f1())
	fmt.Println("f2 result: ", f2())
	fmt.Println("f3 result: ", f3())
	fmt.Println("f4 result: ", f4())
}

/**
原因就是return会将返回值先保存起来，对于无名返回值来说，
保存在一个临时对象中，defer是看不到这个临时对象的；
而对于有名返回值来说，就保存在已命名的变量中。
*/
func f1() int {
	var i int
	fmt.Printf("i: %p \n", &i)
	defer func() {
		i++
		fmt.Printf("i: %p \n", &i)
		fmt.Println("f11: ", i)
	}()

	defer func() {
		i++
		fmt.Printf("i: %p \n", &i)
		fmt.Println("f12: ", i)
	}()

	i = 1000
	return i
}

func f2() (i int) {
	fmt.Printf("i: %p \n", &i)
	defer func() {
		i++
		fmt.Printf("i: %p \n", &i)
		fmt.Println("f21: ", i)
	}()

	defer func() {
		i++
		fmt.Printf("i: %p \n", &i)
		fmt.Println("f22: ", i)
	}()
	return i
}

func f3() (i int) {
	defer func() {
		i = 10
		fmt.Println("defer: ", i)
	}()

	i = 100
	return i
}

func f4() int {
	var i int
	defer func() {
		i = 10
		fmt.Println("defer: ", i)
	}()

	i = 100
	return i
}
