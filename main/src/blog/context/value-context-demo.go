package main

import (
	"context"
	"fmt"
	"time"
)

func HandleRequest2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done")
			return
		default:
			fmt.Println("param value: ", ctx.Value("param"))
			time.Sleep(time.Second * 2)
		}
	}
}

func main() {
	parentCtx, _ := context.WithTimeout(context.Background(), time.Second*5)

	ctx := context.WithValue(parentCtx, "param", "1")
	ctx2 := context.WithValue(ctx, "param2", "2")
	fmt.Println(ctx2.Value("param"))
	fmt.Println(ctx2.Value("param2"))
	go HandleRequest2(ctx)
	time.Sleep(time.Second * 10)
}
