package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	sayHelloFunc := func() {
		defer wg.Done()
		fmt.Println("hello")
	}

	wg.Add(1)
	go sayHelloFunc()
	wg.Wait()

	//
	ch := make(chan string, 3)
	ch <- "1"
	ch <- "2"
	ch <- "3"
	ss := <-ch
	fmt.Println("ch", string(ss))
	ch <- "4"

	sss, ok := <-ch
	fmt.Printf("%v, %v", ok, sss)
	fmt.Println()

	//
	inStream := make(chan int)
	close(inStream)
	interger, ok := <-inStream
	fmt.Printf("%v: %v \n", ok, interger)
	//inStream <- 1	panic: send on closed channel

	intStream2 := make(chan int)
	go func() {
		defer close(intStream2)
		for i := 1; i < 5; i++ {
			intStream2 <- i
		}
	}()

	for interger := range intStream2 {
		fmt.Println(interger)
	}

}
