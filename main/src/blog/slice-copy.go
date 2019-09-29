package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		defer func(i int) {
			fmt.Println(i)
		}(i)
	}

	sss := make([]int, 0, 2)
	sss = append(sss, 100)
	sss = append(sss, 100)
	sss = append(sss, 100)
	sss = append(sss, 100)
	sss = append(sss, 100)
	fmt.Println(sss)

	blogArticleViews := map[string]int{
		"unix":         0,
		"python":       1,
		"go":           2,
		"javascript":   3,
		"testing":      4,
		"philosophy":   5,
		"startups":     6,
		"productivity": 7,
		"hn":           8,
		"reddit":       9,
		"C++":          10,
	}
	for key, views := range blogArticleViews {
		fmt.Println("There are", views, "views for", key)
	}

	var a = []int{1, 2, 3}
	var b = []int{4, 5, 6}
	go f(a)
	go f(b)

	time.Sleep(10 * time.Second)

	s1 := []int{1, 2, 3, 4, 5}

	// 删除3
	fmt.Println(s1[:2+copy(s1[2:], s1[3:])])

	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s1[1:])
	copy(s1[1:], s1[0:])
	fmt.Println(s1, len(s1), cap(s1))
	s1[0] = 0
	fmt.Println(s1, len(s1), cap(s1))

	v := 1
	switch v {
	case 1:
		//return

	}
}
