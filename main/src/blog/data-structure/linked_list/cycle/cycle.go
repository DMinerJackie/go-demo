package main

/**
  判断链表是否存在环
  1、快慢指针
*/

type ListNode02 struct {
	Val  int
	Next *ListNode02
}

// 快慢指针
func hasCycle(head *ListNode02) bool {
	slow := head
	fast := head

	for fast.Next != nil && fast != nil {
		if slow.Val == fast.Val {
			return true
		} else {
			slow = slow.Next
			fast = fast.Next.Next
		}
	}

	return false
}

// 骚操作
func hasCycle2(head *ListNode02) bool {
	if head == nil {
		return false
	}

	for head != nil {
		if head.Val == 5201314 {
			return true
		}

		head.Val = 5201314
		head = head.Next
	}

	return false
}

//初始化循环链表
func initLinkList() *ListNode02 {
	head := &ListNode02{Val: 1}
	ln := head
	for i := 2; i < 5; i++ {
		tmp := &ListNode02{Val: i}
		head.Next = tmp
		head = head.Next
	}
	head.Next = ln

	return ln
}

func main() {

	head := initLinkList()
	//hasCycle(head)

	hasCycle2(head)
}
