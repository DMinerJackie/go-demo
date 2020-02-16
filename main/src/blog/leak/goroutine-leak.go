package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

func sumInt(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

// HTTP handler for /sum
func sumConcurrent2(w http.ResponseWriter, r *http.Request) {
	s := []int{7, 2, 8, -9, 4, 0}

	c1 := make(chan int)
	c2 := make(chan int)

	go sumInt(s[:len(s)/2], c1)
	go sumInt(s[len(s)/2:], c2)

	// 这里故意不在c2中读取数据，导致向c2写数据的协程阻塞。
	x := <-c1

	// write the response.
	fmt.Fprintf(w, strconv.Itoa(x))
}

func main() {
	StasticGroutine := func() {
		for {
			time.Sleep(1e9)
			total := runtime.NumGoroutine()
			fmt.Println(total)
		}
	}

	go StasticGroutine()

	http.HandleFunc("/sum", sumConcurrent2)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
