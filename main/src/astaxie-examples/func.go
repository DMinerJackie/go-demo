package main

import "fmt"

/**
但如果你的函数是导出的(首字母大写)，
官方建议：最好命名返回值，因为不命名返回值，
虽然使得代码更加简洁了，但是会造成生成的文档可读性差。
*/
func SumAndProduct(A, B int) (add int, Multiplied int) {
	add = A + B
	Multiplied = A * B
	return
}

func f1(args ...int) {
	for flag, n := range args {
		fmt.Println(flag, n)
	}
}

func main() {
	SumAndProduct(1, 2)
	f1(1, 2, 3, 4, 5, 6)
}
