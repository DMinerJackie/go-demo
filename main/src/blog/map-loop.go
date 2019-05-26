package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["Jackie"] = "Zheng"
	m["Location"] = "Shanghai"

	for key := range m {
		fmt.Println(key, m[key])
	}

	for key, value := range m {
		fmt.Println(key, value)
	}
}
