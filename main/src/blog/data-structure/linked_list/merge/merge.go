package main

import (
	"fmt"
)

/**
两个有序的链表合并
*/
type ListNode03 struct {
	Val  int
	Next *ListNode03
}

func mergeTwoLists(l1 *ListNode03, l2 *ListNode03) (result *ListNode03) {
	if l1 == nil && l2 == nil {
		return nil
	}

	result = &ListNode03{}
	cur := result
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = &ListNode03{Val: l1.Val}
			l1 = l1.Next
		} else {
			cur.Next = &ListNode03{Val: l2.Val}
			l2 = l2.Next
		}

		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	}

	if l2 != nil {
		cur.Next = l2
	}
	return result.Next // 这里将初始化的头结点过滤掉
}

func initLinkList() *ListNode03 {
	head := &ListNode03{Val: 0}
	ln := head
	for i := 2; i < 10; i = i + 2 {
		tmp := &ListNode03{Val: i}
		head.Next = tmp
		head = head.Next
	}

	return ln
}

func initLinkList2() *ListNode03 {
	head := &ListNode03{Val: 1}
	ln := head
	for i := 3; i < 10; i = i + 2 {
		tmp := &ListNode03{Val: i}
		head.Next = tmp
		head = head.Next
	}

	return ln
}

type st struct {
	Name string
}

func main() {

	l1 := initLinkList()
	l2 := initLinkList2()

	result := mergeTwoLists(l1, l2)
	for result != nil {
		fmt.Print(result.Val, " ")
		result = result.Next
	}
}
