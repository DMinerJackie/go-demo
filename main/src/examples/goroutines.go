package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(time.Duration(2) * time.Second)
	}
}

func main() {
	f("direct")

	go f("goroutines")

	// 通过在f函数中添加 time.Sleep(time.Duration(2)*time.Second)，可以看出上面异步和下面匿名异步执行的顺序是同时进行的
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Scanln()
	fmt.Println("done")
}
