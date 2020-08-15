package main

import (
	"container/heap"
	"fmt"
)

type myHeap []int

/* 实现排序 */
func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// 最小堆实现
func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

/* 实现往堆中添加元素 */
func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

/* 实现删除堆中元素 */
func (h *myHeap) Pop() (v interface{}) {
	*h, v = (*h)[:len(*h)-1], (*h)[len(*h)-1]
	return
}

// 按层来遍历和打印堆数据，第一行只有一个元素，即堆顶元素
func (h myHeap) printHeap() {
	n := 1
	levelCount := 1
	for n <= h.Len() {
		fmt.Println(h[n-1 : n-1+levelCount])
		n += levelCount
		levelCount *= 2
	}
}

func main() {
	data := [7]int{13, 12, 45, 23, 11, 9, 20}
	aHeap := new(myHeap)
	for i := 0; i < len(data); i++ {
		aHeap.Push(data[i])
	}
	aHeap.printHeap()

	// 堆排序处理
	heap.Init(aHeap)
	aHeap.printHeap()
}
