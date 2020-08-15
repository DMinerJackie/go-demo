package main

import "fmt"

type ListNode01 struct {
	Val  int
	Next *ListNode01
}

func isPalindrome(head *ListNode01) bool {
	if head == nil || head.Next == nil {
		return true
	}

	fast := head
	slow := head

	// 快慢指针，快指针每次移动两步，慢指针每次移动一步
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	//如果是奇数，则慢指针再移动一步，达到中心点
	if fast != nil {
		slow = slow.Next
	}

	// 慢指针反转
	var pre *ListNode01 = nil
	for slow != nil {
		pre, slow, slow.Next = slow, slow.Next, pre
	}

	cur := head
	for cur != nil && pre != nil {
		if cur.Val == pre.Val {
			cur = cur.Next
			pre = pre.Next
		} else {
			return false
		}
	}

	//if cur != nil && cur.Next != nil {
	//	return false
	//} else {
	//	return true
	//}

	return true
}

func main() {
	head := new(ListNode01)
	head.Val = 1
	//head := ln
	ln2 := new(ListNode01)
	ln2.Val = 2
	ln3 := new(ListNode01)
	ln3.Val = 3
	ln4 := new(ListNode01)
	ln4.Val = 3
	ln5 := new(ListNode01)
	ln5.Val = 2
	ln6 := new(ListNode01)
	ln6.Val = 1
	head.Next = ln2
	ln2.Next = ln3
	ln3.Next = ln4
	ln4.Next = ln5
	ln5.Next = ln6

	//for i := 1; i < 5; i++ {
	//	tmp := &ListNode01{Val:i}
	//	head.Next = tmp
	//	head = head.Next
	//}
	fmt.Println(isPalindrome(head))
}
