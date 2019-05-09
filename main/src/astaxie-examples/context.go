package main

import (
	"context"
	"log"
	"os"
	"time"
)

var logg *log.Logger

func someHandler() {
	ctx, cancel := context.WithCancel(context.Background())
	go doStuff(ctx)

	//10秒后取消doStuff
	time.Sleep(10 * time.Second)
	cancel()

}

//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
func doStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			logg.Printf("done")
			return
		default:
			logg.Printf("work")
		}
	}
}

func timeoutHandler() {
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	// go doTimeOutStuff(ctx)
	go doStuff(ctx)

	time.Sleep(10 * time.Second)

	cancel()

}

func main() {
	logg = log.New(os.Stdout, "", log.Ltime)
	timeoutHandler()
	logg.Printf("end")
}

//func main() {
//	logg = log.New(os.Stdout, "", log.Ltime)
//	someHandler()
//	logg.Printf("down")
//}

//// 通过context.WithValue传值
//func main() {
//	ctx, cancel := context.WithCancel(context.Background())
//
//	valueCtx := context.WithValue(ctx, "kkk", "add value")
//
//	go watch(valueCtx)
//	time.Sleep(10 * time.Second)
//	cancel()
//
//	time.Sleep(5 * time.Second)
//}
//
//func watch(ctx context.Context) {
//	for {
//		select {
//		case <-ctx.Done():
//			//get value
//			fmt.Println(ctx.Value("kkk"), "is cancel")
//
//			return
//		default:
//			//get value
//			fmt.Println(ctx.Value("kkk"), "int goroutine")
//
//			time.Sleep(2 * time.Second)
//		}
//	}
//}

// 超时取消 context.WithTimeout
//import (
//	"fmt"
//	"sync"
//	"time"
//
//	"golang.org/x/net/context"
//)
//
//var (
//	wg sync.WaitGroup
//)
//
//func work(ctx context.Context) error {
//	defer wg.Done()
//
//	for i := 0; i < 1000; i++ {
//		select {
//		case <-time.After(2 * time.Second):
//			fmt.Println("Doing some work ", i)
//
//		// we received the signal of cancelation in this channel
//		case <-ctx.Done():
//			fmt.Println("Cancel the context ", i)
//			return ctx.Err()
//		}
//	}
//	return nil
//}
//
//func main() {
//	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
//	defer cancel()
//
//	fmt.Println("Hey, I'm going to do some work")
//
//	wg.Add(1)
//	go work(ctx)
//	wg.Wait()
//
//	fmt.Println("Finished. I'm going home")
//}

//// 截止取消 context.withDeadline
//import (
//	"context"
//	"fmt"
//	"time"
//)
//
//func main() {
//	d := time.Now().Add(1 * time.Second)
//	ctx, cancel := context.WithDeadline(context.Background(), d)
//
//	// Even though ctx will be expired, it is good practice to call its
//	// cancelation function in any case. Failure to do so may keep the
//	// context and its parent alive longer than necessary.
//	defer cancel()
//
//	select {
//	case <-time.After(2 * time.Second):
//		fmt.Println("oversleep")
//	case <-ctx.Done():
//		fmt.Println(ctx.Err())
//	}
//}
