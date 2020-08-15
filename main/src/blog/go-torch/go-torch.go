package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {

	go func() {
		http.ListenAndServe(":6789", nil)
	}()

	time.Sleep(time.Minute * 10)

}
