/**
多个senders，一个receivers
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	const Max = 10000
	const NumSenders = 1000

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})

	for i := 0; i < NumSenders; i++ {
		go func() {
			select {
			case <-stopCh:
				return
			case dataCh <- rand.Intn(Max):
			}
		}()
	}

	go func() {
		for value := range dataCh {
			if value == Max-1 {
				fmt.Println("send stop signal to senders.", value)
				close(stopCh)
				return
			}
			fmt.Println(value)
		}
	}()

	select {
	case <-time.After(time.Second * 10):
	}
}
