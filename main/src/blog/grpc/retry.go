package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	//s := "abcd"
	//for _, v := range s {
	//	fmt.Printf("%c", v)
	//}
	//m := make(map[string]string)
	//fmt.Println(m["hh"])
	//
	//var bs bytes.Buffer
	//bs.WriteString("jasldfjlasjdf12221")
	//bs.WriteString("ababababab")
	//r := md5.Sum(bs.Bytes())
	//value, tErr := strconv.ParseInt(hex.EncodeToString(r[0:7]), 16, 64)
	//fmt.Println(value, tErr)
	//value1, _ := strconv.ParseInt(hex.EncodeToString(r[0:8]), 16, 64)
	//fmt.Println(value1, tErr)
	//
	//fmt.Println(math.MaxInt64)
	//
	//flag := true
	//for flag {
	//	if 1== 1 {
	//		break
	//		fmt.Println("break")
	//	}
	//	fmt.Println("out")
	//}
	//
	//var g errgroup.Group
	//var urls = []string{
	//	"http://www.baidu.com/",
	//	"http://www.baidu.com/",
	//	"http://www.baidu.com/",
	//}
	//for _, url := range urls {
	//	// Launch a goroutine to fetch the URL.
	//	url := url // https://golang.org/doc/faq#closures_and_goroutines
	//	g.Go(func() error {
	//		// Fetch the URL.
	//		resp, err := http.Get(url)
	//		if err == nil {
	//			resp.Body.Close()
	//		}
	//		return err
	//	})
	//}
	//// Wait for all HTTP fetches to complete.
	//if err := g.Wait(); err == nil {
	//	fmt.Println("Successfully fetched all URLs.")
	//} else {
	//	fmt.Printf("error: %v", err)
	//}

	var waitTime time.Duration
	waitTime = time.Duration(time.Second * 5)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*4))
	cancel()
	defer cancel()
	timer := time.NewTimer(waitTime)
	select {
	case <-ctx.Done():
		//timer.Stop()
		timer.Reset(time.Second)
		fmt.Println("ctx done")
	case <-timer.C:
		fmt.Println("timeout")
	}

	select {
	case <-timer.C:
		fmt.Println("timeout finally")
	}
}
