package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func init() {
	fmt.Println("hhhhhh")
}

func f1(arr *[]int64) {
	(*arr)[0] = 11
	(*arr)[4] = 11
	(*arr)[5] = 11
}

func f2(arr []int64) {
	arr[0] = 11
	arr[4] = 11
	arr[5] = 11
	arr = append(arr, 11)
	arr = append(arr, 11)
	arr = append(arr, 11)
	arr = append(arr, 11)
	arr = append(arr, 11)
}

func main() {
	s1 := make([]int64, 6, 6)
	s1[0] = 1
	s1[1] = 2
	s1[2] = 3
	fmt.Println(s1)
	f1(&s1)
	fmt.Println(s1)

	s2 := make([]int64, 6, 6)
	s2[0] = 1
	s2[1] = 2
	s2[2] = 3
	fmt.Println(s2)
	f2(s2)
	fmt.Println(s2)

	fmt.Println(time.Now().Unix())
	fmt.Println(1 << 10)
	var line []byte
	reader := bufio.NewReader(strings.NewReader("a b c d e f g"))
	line, _ = reader.ReadSlice(' ')
	fmt.Printf("the line:%s\n", line)
	// 这里可以换上任意的 bufio 的 Read/Write 操作
	line, _ = reader.ReadSlice(' ')
	fmt.Printf("the line:%s\n", line)
	line, _ = reader.Peek(7) // 只是获取不会取出
	fmt.Println(string(line))
	line, _ = reader.ReadSlice(' ')
	fmt.Printf("the line:%s\n", line)

	//scanner := bufio.NewScanner(os.Stdin)
	//for scanner.Scan() {
	//	fmt.Println(scanner.Text()) // Println will add back the final '\n'
	//}
	//if err := scanner.Err(); err != nil {
	//	fmt.Fprintln(os.Stderr, "reading standard input:", err)
	//}

	file, err := os.Create("scanner.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("http://studygolang.com.\nIt is the home of gophers.\nIf you are studying golang, welcome you!")
	// 将文件 offset 设置到文件开头
	file.Seek(0, os.SEEK_SET)
	scanner2 := bufio.NewScanner(file)
	for scanner2.Scan() {
		fmt.Println(scanner2.Text())
	}
}
