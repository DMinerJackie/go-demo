package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func main() {
	c := cron.New()
	c.AddFunc("*/3 * * * * *", func() {
		fmt.Println("every 3 seconds executing")
	})

	go c.Start()
	defer c.Stop()

	select {
	case <-time.After(time.Second * 10):
		return
	}
}
