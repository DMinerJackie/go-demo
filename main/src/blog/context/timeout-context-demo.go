package main

import (
	"context"
	"fmt"
	"time"
)

func HandleRequest(ctx context.Context) {
	go WriteRedis2(ctx)
	go WriteDatabase2(ctx)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandleRequest Done")
			return
		default:
			fmt.Println("HandleRequest running")
			time.Sleep(2 * time.Second)
		}
	}

}

func WriteRedis2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Write Redis done")
			return
		default:
			fmt.Println("Write Redis doing")
			time.Sleep(2 * time.Second)
		}
	}
}

func WriteDatabase2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Write Database done")
			return
		default:
			fmt.Println("Write Database doing")
			time.Sleep(2 * time.Second)

		}
	}
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	go HandleRequest(ctx)

	time.Sleep(time.Second * 10)
}
