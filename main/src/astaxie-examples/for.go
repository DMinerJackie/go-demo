package main

import "fmt"

func main() {
	for i, j := 1, 1; i < 100; i++ {
		fmt.Println(i, j)
	}

	/**
	类似while的用法
	*/
	sum := 1
	for sum < 100 {
		sum += sum
		fmt.Println(sum)
	}

	map1 := make(map[string]string)
	map1["aa"] = "aav"
	for k, v := range map1 {
		fmt.Println("key:", k)
		fmt.Println("value:", v)
	}

	map2 := map[string]string{"bb": "bbv"}
	for k, v := range map2 {
		fmt.Println("key:", k)
		fmt.Println("value:", v)
	}
}
