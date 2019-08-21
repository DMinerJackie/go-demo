/**
多个senders，多个receivers
*/
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	const Max = 10000
	const NumSenders = 1000
	const NumReceivers = 10

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})

	toStop := make(chan string, 1)

	var stoppedBy string

	go func() {
		stoppedBy = <-toStop
		close(stopCh)
	}()

	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(Max)
				if value == 0 {
					select {
					case toStop <- "sender#" + id:
					default:
					}
					return
				}

				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			for {
				select {
				case <-stopCh:
					return
				case value := <-dataCh:
					if value == Max-1 {
						select {
						case toStop <- "receive#" + id:
						default:

						}
						return
					}
					fmt.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}

	select {
	case <-time.After(time.Second * 10):
	}
}
