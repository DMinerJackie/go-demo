package main

import "fmt"

/*
	快速排序
	快速排序 复杂度 N * (logN)
    原理为有两个指针:low, high 分别指向列表的第一个元素,和列表最后一个元素,和一个中间值 mid(就是取列表的第一个元素)
    一开始 low 默认情况下就指向了mid, high 指向列表最后一个元素
    [8, 3, 15, 7, 6, 2]
     |               |
    low             high
    如果high 指向的值比mid大,则high往前移动,如果high比mid 小,就把high 指向的值和low交换
    lst[low], lst[high] = lst[high], lst[low]
    如果low比mid小low往前移动一位, 如果low比mid大,则low指向的值和high 指向的值交换
    lst[high], lst[low] = lst[low], lst[high]
    当low 和 high 指针移动到相同的位置时,则mid就应该放在这个位置
*/

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}

}

func main() {

	s := "a我bc123"
	for _, e := range s {
		fmt.Printf("%T, %b, %d\n", e, e, e)
	}
	fmt.Println("====")
	bb := []byte(s)
	for _, b := range bb {
		fmt.Printf("%T, %b, %d\n", b, b, b)
	}

	list := []int{22, 56, 38, 101, 1, 18, 20, 30}
	fmt.Println("排序前", list)
	quick(list, 0, len(list)-1)
	fmt.Println("排序后", list)

}

func quick(lst []int, start, end int) {
	if start >= end {
		return
	}
	low := start
	high := end
	mid := lst[start]

	for low < high {
		for low < high && lst[high] >= mid {
			high -= 1
		}
		//出了循环说明 high 小于等于 mid 则需要交换位置
		lst[high], lst[low] = lst[low], lst[high]

		for low < high && lst[low] < mid {
			low += 1
		}
		// 出了循环说明 low 大于等于mid 则需要交换位置
		lst[low], lst[high] = lst[high], lst[low]
	}
	// 大的循环退出以后就找到了 mid的位置,此时mid 左边的都小于mid, 右边都大于mid,以 mid 为界限再排序左右两边
	lst[low] = mid
	quick(lst, start, low-1) // 排序左边的
	quick(lst, low+1, end)   // 排序右边的
}
