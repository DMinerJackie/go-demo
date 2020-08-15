package main

/**
删除链表倒数第n个节点
*/
type ListNode04 struct {
	Val  int
	Next *ListNode04
}

func removeNthFromEnd(head *ListNode04, n int) *ListNode04 {
	sentinel := &ListNode04{Next: head}
	fast, slow, index := sentinel, sentinel, 0
	for index < n+1 { // 快指针先走n步
		fast = fast.Next
		index++
	}
	for fast != nil { // 快慢指针同步走，直到快指针到达链表最后一个节点
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return sentinel.Next
}

//func removeNthFromEnd(head *ListNode04, n int) *ListNode04 {
//	var i int
//	tmp := head
//	targetNode := head
//	for tmp.Next != nil {
//		if i >= n {
//			targetNode = targetNode.Next
//		}
//		tmp = tmp.Next
//		i++ // 计数器
//	}
//
//	targetNode.Next = targetNode.Next.Next
//
//	return head
//}

func initLinkList() *ListNode04 {
	head := &ListNode04{Val: 1}
	ln := head
	for i := 2; i < 6; i++ {
		tmp := &ListNode04{Val: i}
		head.Next = tmp
		head = head.Next
	}

	return ln
}

func main() {
	head := initLinkList()
	removeNthFromEnd(head, 1)
}
