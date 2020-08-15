package main

import (
	"crypto/md5"
	"fmt"
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	//sort.Slice(people, func(i, j int) bool {
	//	return (people[i][0] > people[j][0]) || (people[i][0] == people[j][0] && people[i][1] < people[j][1])
	//})
	//
	//for i, p := range people {
	//	copy(people[p[1]+1:i+1], people[p[1]:i+1])
	//	people[p[1]] = p
	//}
	//
	//return people

	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}

		return people[i][0] > people[j][0]
	})

	res := make([][]int, 0)

	for i := 0; i < len(people); i++ {
		p := people[i][1]
		if p >= len(res) {
			res = append(res, people[i])
		} else {
			res = append(res, res[len(res)-1])
			copy(res[p+1:], res[p:])
			res[p] = people[i]
		}
	}

	return people
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}

	fmt.Println(slice1[0:1])
	fmt.Println(slice1[1:1])
	//copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置

	s := fmt.Sprintf("%x", md5.Sum([]byte("你好你好你好你好你好你好你好你好你好你好")))
	fmt.Println(s)
	fmt.Println("tu_b_111004142_1594803290081255258_" + s)
	fmt.Println(len("tu_b_111004142_1594803290081255258_dbeec291ac7d07607a2fee29cf24b8f3"))

	people := [][]int{[]int{7, 0}, []int{4, 4}, []int{7, 1}, []int{5, 0}, []int{6, 1}, []int{5, 2}}
	reconstructQueue(people)
}
