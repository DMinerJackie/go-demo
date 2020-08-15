package cocurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestC(t *testing.T) {
	i := -1
	fmt.Println([]int64{}[i])
	ch := make(chan int, 5)

	var wg sync.WaitGroup

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(<-ch)
		}()
	}

	time.Sleep(time.Second)
	wg.Wait()
	//close(ch)

	time.Sleep(time.Second * 10)
}
