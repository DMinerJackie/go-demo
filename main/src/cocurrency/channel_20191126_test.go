package cocurrency

import (
	"fmt"
	"testing"
	"time"
)

func Test20191126Channel(t *testing.T) {

	s := []int{1, 10, 100, 22}
	for index, se := range s {
		se = se + 1
		fmt.Println(index, ":", se)
	}

	for index, se := range s {
		se = se + 1
		fmt.Println(index, ":", se)
	}
	fmt.Println(time.Until(time.Unix(1575343970, 0)).Hours())
	fmt.Println(time.Unix(1575343970, 0).Sub(time.Now()).Hours())

	queue := make(chan int, 10)
	go func() {
		queue <- 10
		queue <- 20

	}()

	//close(queue)

	time.Sleep(time.Second)

	//close(queue)

	for q := range queue {
		fmt.Println(q)
	}

	//fmt.Println(<-queue)
	//fmt.Println(<-queue)
	//fmt.Println(<-queue)

	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	time.Sleep(time.Millisecond * 200)

	close(ch)

	//go func() {
	for i := 0; i < 14; i++ {
		go func(i int) {
			fmt.Println("i: ", i, "go: ", <-ch)
		}(i)
		//fmt.Println(<-ch)

	}
	//}()

	time.Sleep(time.Second)
}
