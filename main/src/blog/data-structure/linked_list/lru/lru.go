package lru

import "fmt"

type LinkNode struct {
	Value int
	Next  *LinkNode
}

/**
使用单链表方式自己简单实现了LRU
1.初始化
2.访问
3.插入（已满和未满）
*/

// 初始化
func initData() *LinkNode {
	ln := &LinkNode{}
	head := ln
	for i := 1; i < 5; i++ {
		tmp := &LinkNode{Value: i}
		head.Next = tmp
		head = head.Next
	}
	return ln
}

// 打印
func printData(ln *LinkNode) {
	for ln != nil {
		fmt.Print(ln.Value, " ")
		ln = ln.Next
	}
}

// 获取指定值
func get(ln *LinkNode, expectValue int) *LinkNode {
	if ln == nil {
		return nil
	}

	times := 0

	node := ln
	if node.Value == expectValue {
		return ln
	}
	for node != nil && node.Next != nil {
		if node.Next.Value == expectValue {
			node.Next = node.Next.Next
			//ln.Next = node
			newNode := &LinkNode{Value: expectValue}
			newNode.Next = ln
			//node.Next.Next = ln
			//ln.Next = node
			return newNode
		}

		times++
		node = node.Next
		if times == 5 { // 链表已满
			node.Next = nil
			return ln
		}
	}

	// 如果链表中没有找到期望值
	if times < 6 { // 链表未满
		newNode := &LinkNode{Value: expectValue}
		newNode.Next = ln
		return newNode
	}
	return ln
}

func main() {
	ln := initData()
	printData(ln)
	fmt.Println("========")
	getLn := get(ln, 3)
	printData(getLn)
	fmt.Println("========")
	get2Ln := get(getLn, 2)
	printData(get2Ln)
}
