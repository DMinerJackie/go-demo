package main

import (
	"fmt"
	"time"
)

var (
	domainSyncChan = make(chan int, 10)
)

func main() {
	//num := runtime.NumCPU()
	//runtime.GOMAXPROCS(num)
	//fmt.Println(num)

	for i := 0; i < 10; i++ {
		domainName := i
		go domainPut(domainName)
	}
	time.Sleep(time.Second * 2)
}

func domainPut(num int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("error to chan put. %v", err)
			fmt.Println()
		}
	}()

	domainSyncChan <- num
	if num == 7 {
		panic("error......")
	}
}
