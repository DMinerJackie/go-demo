package main

import (
	"context"
	custom_pool "go-demo/main/src/custom-pool"
)

func main() {

	ctx := context.Background()
	config := &custom_pool.Config{
		MaxConn: 1,
		MaxIdle: 1,
	}
	conn := custom_pool.Prepare(ctx, config)
	if _, err := conn.New(ctx); err != nil {
		return
	}
	if _, err := conn.New(ctx); err != nil {
		return
	}
	//go conn.Release(ctx)
	if _, err := conn.New(ctx); err != nil {
		return
	}
	go conn.Release(ctx)
	if _, err := conn.New(ctx); err != nil {
		return
	}
	go conn.Release(ctx)
	if _, err := conn.New(ctx); err != nil {
		return
	}
	go conn.Release(ctx)
	if _, err := conn.New(ctx); err != nil {
		return
	}
	go conn.Release(ctx)
	if _, err := conn.New(ctx); err != nil {
		return
	}

}
