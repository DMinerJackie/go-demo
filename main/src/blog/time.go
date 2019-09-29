package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println(time.Now())
	//
	//timer := time.NewTimer(time.Second * 5)
	//
	//<-timer.C
	//
	//fmt.Println(time.Now())

	//t, _ := time.Parse("20060102", "20180531")
	//fmt.Println(t)

	//loc, _ := time.LoadLocation("PRC")
	//t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-05-31 09:22:19", loc)
	//fmt.Println(t)
	//fmt.Println(t.Format("2006-01-02 15:04:05"))
	//fmt.Println(t.String())
	//
	//seconds := 100000
	//fmt.Println(time.Duration(seconds) * time.Second)

	//wait := sync.WaitGroup{}
	//fmt.Println("start", time.Now())
	//
	//wait.Add(1)
	//
	//timer := time.AfterFunc(time.Second * 3, func() {
	//	fmt.Println("get timer", time.Now())
	//	wait.Done()
	//})
	//
	//time.Sleep(time.Second)
	//fmt.Println("sleep", time.Now())
	//
	//timer.Reset(time.Second * 2)
	//
	//wait.Wait()

	//fmt.Println("start", time.Now())
	//
	//ticker := time.NewTicker(time.Second)
	//
	//go func() {
	//	for tick := range ticker.C {
	//		fmt.Println("tick at", tick)
	//	}
	//}()
	//
	//time.Sleep(time.Second * 5)
	//ticker.Stop()
	//
	//fmt.Println("stoped", time.Now())

	//fmt.Println("start", time.Now())
	//
	//timer := time.After(time.Second)
	//
	//select {
	//case t := <-timer:
	//	fmt.Println("get timer", t)
	//}
	//
	//fmt.Println("stoped", time.Now())

	var t time.Time // 定义 time.Time 类型变量
	t = time.Now()  // 获取当前时间
	fmt.Printf("时间: %v, 时区:  %v,  时间类型: %T\n", t, t.Location(), t)
	// 时间: 2017-02-22 09:06:05.816187261 +0800 CST, 时区:  Local,  时间类型: time.Time

	// time.UTC() time 返回UTC 时区的时间
	fmt.Printf("时间: %v, 时区:  %v,  时间类型: %T\n", t.UTC(), t.UTC().Location(), t)
	// 时间: 2017-02-22 01:07:15.179280004 +0000 UTC, 时区:  UTC,  时间类型: time.Time

	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := data[2:4:9] // data[low, high, max]
	fmt.Println(slice)

}
