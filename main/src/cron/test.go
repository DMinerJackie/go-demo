//package main
//
//import "fmt"
//
//type rect struct {
//	width, height int
//}
//
//func (r *rect) area() int {
//	return r.width * r.height
//}
//
//func (r rect) perim() int {
//	return 2*r.width + 2*r.height
//}
//
//func main() {
//	r := rect{width: 10, height: 20}
//	fmt.Println("area:", r.area())
//	fmt.Println("perim:", r.perim())
//
//	rp := &r
//	fmt.Println("area:", rp.area())
//	fmt.Println("perim:", rp.perim())
//}
//
//
////5、定义函数类型
////以下是定义一个函数类型handler
//type handler func (name string) int
//
////针对这个函数类型可以再定义方法，如：
//func (h handler) add(name string) int {
//	return h(name) + 10
//}

//下面让我们详细看一下例子，其中涉及了函数、函数的方法、结构体方法、接口的使用。
package main

import (
	"fmt"
)

//定义接口
type adder interface {
	add(string) int
}

//定义函数类型
type handler func(name string) int

//实现函数类型方法
func (h handler) add(name string) int {
	return h(name) + 10
}

//函数参数类型接受实现了adder接口的对象（函数或结构体）
func process(a adder) {
	fmt.Println("process:", a.add("taozs"))
}

//另一个函数定义
func doubler(name string) int {
	return len(name) * 2
}

//非函数类型
type myint int

//实现了adder接口
func (i myint) add(name string) int {
	return len(name) + int(i)
}

func main() {
	//注意要成为函数对象必须显式定义handler类型
	var my handler = func(name string) int {
		return len(name)
	}
	//以下是函数或函数方法的调用
	fmt.Println(my("taozs")) //调用函数

	fmt.Println(my.add("taozs")) //调用函数对象的方法

	fmt.Println(handler(doubler).add("taozs")) //doubler函数显式转换成handler函数对象然后调用对象的add方法

	//以下是针对接口adder的调用
	process(my) //process函数需要adder接口类型参数

	process(handler(doubler)) //因为process接受的参数类型是handler，所以这儿要强制转换

	process(myint(8)) //实现adder接口不仅可以是函数也可以是结构体
}
