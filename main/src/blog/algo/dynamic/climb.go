package main

import (
	"fmt"
	"time"
)

func climb1(i, n int) int {
	if i == n {
		return 1
	}

	if i > n {
		return 0
	}

	return climb1(i+1, n) + climb1(i+2, n)
}

const (
	info  = "m"
	info2 = "m2"
	b1    = iota
	b2    = iota
)

func climb2(n int) int {
	var tempMap = make(map[int]int)
	tempMap[1] = 1
	tempMap[2] = 2
	for i := 3; i <= n; i++ {
		tempMap[i] = tempMap[i-1] + tempMap[i-2]
	}

	return tempMap[n]
}

func A() int {
	time.Sleep(100 * time.Millisecond)
	return 1
}

func B() int {
	time.Sleep(1000 * time.Millisecond)
	return 2
}

func watShadowDefer(i int) (ret int) {
	ret = i * 2
	if ret > 10 {
		ret := 10
		defer func() {
			ret = ret + 1
		}()
	}
	return
}

func main() {
	fmt.Println(watShadowDefer(50))
	ch := make(chan int, 1)
	go func() {
		select {
		case ch <- A():
		case ch <- B():
		default:
			ch <- 3
		}
	}()

	fmt.Println(<-ch)

	//fmt.Println(b1, b2)
	//
	//fmt.Print(climb2(4))
	//fmt.Println(climb1(1, 5))
}
