package main

import (
	"fmt"
	"os"
)

var user = os.Getenv("USER")

func main() {
	//s := "hello"
	//c := []byte(s)  // 将字符串 s 转换为 []byte 类型
	//c[0] = 'c'
	//s2 := string(c)  // 再转换回 string 类型
	//fmt.Printf("%s",s2)

	m1 := map[int]string{}
	fmt.Println("len m1:", len(m1))
	fmt.Println("*", m1[1], "*")
	if v, ok := m1[1]; ok {
		fmt.Print("*", v, "*")
	} else {
		fmt.Print("not exist")
	}
	m1[1] = ""
	if v, ok := m1[1]; ok {
		fmt.Print("*", v, "*")
	} else {
		fmt.Print("not exist")
	}
	fmt.Println("*", m1[1], "*")
	fmt.Println("len m1:", len(m1))

	s1 := "string"
	fmt.Println("s1:", s1)
	//s1[0] = "c"
	s1 = "helloworld"
	fmt.Println("s1: ", s1)
	s := "hello"
	s = "c" + s[1:] // 字符串虽不能更改，但可进行切片操作
	fmt.Printf("%s\n", s)

	//if user != "" {
	//	panic("no value for $USER")
	//	fmt.Println("null")
	//}
}

func throwsPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f() //执行函数f，如果f中出现了panic，那么就可以恢复回来
	return
}
