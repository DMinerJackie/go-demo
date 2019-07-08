package main

import (
	"fmt"
	"time"
)

//func main() {
//	c:=make(chan int)
//	go func() {
//		time.Sleep(time.Second)
//		fmt.Println(<-c)
//		fmt.Println("====")
//		time.Sleep(time.Second)
//	}()
//	c<-666
//	//time.Sleep(time.Second * 2)
//}

func main() {
	go func() {
		chs := make(chan string, 2)
		chs <- "first"
		chs <- "second"

		for ch := range chs {
			fmt.Println(ch)
			if len(chs) == 0 {
				break
			}
		}
	}()

	time.Sleep(time.Second * 2)
}
