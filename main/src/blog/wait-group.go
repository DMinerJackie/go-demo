/**
waitGroup使用
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func calc1(w *sync.WaitGroup, i int) {
	fmt.Println("calc: ", i)
	time.Sleep(time.Second * 2)
	w.Done()
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go calc1(&wg, i)
	}
	wg.Wait()
	fmt.Println("all goroutine finish")
}
