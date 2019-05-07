package main

import "fmt"

func main() {
	var arr1 = [3]int{1, 2, 3}
	var arr2 = arr1
	arr2[1] = 10
	fmt.Println(arr1, arr2)

	var s1 = []int{1, 2, 3, 4, 5}
	var s2 = s1
	s2[2] = 111
	fmt.Println(s1, s2)

	var m1 = map[string]string{"k1": "v1", "k2": "v2", "k3": "v3"}
	var m2 = m1
	m2["k2"] = "vvvvvv"
	fmt.Println(m1, m2)
}
