package main

import "fmt"

//func main()  {
//	//var userChan chan interface{}
//	//userChan = make(chan interface{}, 10)
//
//	//var readOnlyChan <-chan int		// 只读
//	//var writeOnlyChan chan<- int	// 只写
//
//	//userChan <- "nick"
//	//name := <- userChan
//	//name1, ok := <- userChan
//	//fmt.Println(name, name1, ok)
//	//
//	//intChan := make(chan int, 1)
//	//intChan <- 9
//	//close(intChan)
//	//
//	//intChan2 := make(chan int, 10)
//	//for i := 0; i < 100; i++ {
//	//	intChan2 <- i
//	//}
//	//close(intChan2)
//	//
//	//for v := range intChan2 {
//	//	fmt.Println(v)
//	//}
//
//	//userChan := make(chan interface{})
//	//go func() {
//	//	userChan <- "nick"
//	//}()
//	//fmt.Println(<-userChan)
//
//	//userChan := make(chan interface{})
//	//go func() {
//	//	for {
//	//		userChan <- "nick"
//	//	}
//	//}()
//	//for {
//	//	name := <- userChan
//	//	fmt.Println(name)
//	//	time.Sleep(time.Millisecond)
//	//}
//
//
//	intChan := make(chan int, 10)
//
//	for i := 0; i < 10; i++ {
//		intChan <- i
//	}
//	close(intChan)
//	for {
//		//十次后 fatal error: all goroutines are asleep - deadlock!
//		i := <- intChan
//		fmt.Println(i)
//		time.Sleep(time.Second)
//	}
//
//
//}

//func sendData(ch chan<- string) {
//	ch <- "go"
//	ch <- "java"
//	ch <- "c"
//	ch <- "c++"
//	ch <- "python"
//	close(ch)
//}
//
//func getData(ch <-chan string, chColse chan bool) {
//	for {
//		str, ok := <-ch
//		if !ok {
//			fmt.Println("chan is close.")
//			break
//		}
//		fmt.Println(str)
//	}
//	chColse <- true
//}

//func main() {
//	//ch := make(chan string, 10)
//	//chColse := make(chan bool, 1)
//	//go sendData(ch)
//	//go getData(ch, chColse)
//	//<-chColse
//	//close(chColse)
//
//	//ch := make(chan string, 3)
//	//	//ch <- "A"
//	//	//ch <- "B"
//	//	//ch <- "C"
//	//	//ch <- "D"
//
//}

//var s []int
////var lock2 sync.Mutex //互斥锁
//func appendValue(i int) {
//	//lock2.Lock() //加锁
//	s = append(s, i)
//	//lock2.Unlock() //解锁
//}
//
//func main() {
//	for i := 0; i < 10000; i++ {
//		go appendValue(i)
//	}
//	//sort.Ints(s) //给切片排序,先排完序再打印,和下面一句效果相同
//	time.Sleep(time.Second)  //间隔1s再打印,防止一边插入数据一边打印时数据乱序
//	for i, v := range s {
//		fmt.Println(i, ":", v)
//	}
//}

//func main() {
//	m := make(map[int]int)
//	go func() {				//开一个协程写map
//		for i := 0; i < 10000; i++ {
//			m[i] = i
//		}
//	}()
//
//	go func() {				//开一个协程读map
//		for i := 0; i < 10000; i++ {
//			fmt.Println(m[i])
//		}
//	}()
//
//	//time.Sleep(time.Second * 20)
//	for {
//		;
//	}
//}

//var ch1 chan int = make(chan int)
//var ch2 chan int = make(chan int)
//
//func say(s string) {
//	fmt.Println(s)
//	ch1 <- <- ch2 // ch1 等待 ch2流出的数据
//}
//
//func main() {
//	go say("hello")
//	<- ch1  // 堵塞主线
//}

//func main() {
//fmt.Println("run in main goroutine")
//	//go func() {
//	//	fmt.Println("run in child goroutine")
//	//	go func() {
//	//		fmt.Println("run in grand child goroutine")
//	//		go func() {
//	//			fmt.Println("run in grand grand child goroutine")
//	//		}()
//	//	}()
//	//}()
//	//time.Sleep(time.Second)
//	//fmt.Println("main goroutine will quit")

//}

var ch chan int = make(chan int)

func foo(id int) { //id: 这个routine的标号
	ch <- id
}

func main() {
	// 开启5个routine
	for i := 0; i < 5; i++ {
		go foo(i)
	}

	// 取出信道中的数据
	for i := 0; i < 5; i++ {
		fmt.Print(<-ch)
	}
}
