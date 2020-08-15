package main

type ListNode05 struct {
	Val  int
	Next *ListNode05
}

func middleNode(head *ListNode05) *ListNode05 {
	sentinel := head
	slow, fast := sentinel, sentinel
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast.Next != nil {
		slow = slow.Next
	}

	return slow
}

func initLinkList() *ListNode05 {
	head := &ListNode05{Val: 1}
	ln := head
	for i := 2; i < 6; i++ {
		tmp := &ListNode05{Val: i}
		head.Next = tmp
		head = head.Next
	}

	return ln
}

func main() {

	head := initLinkList()
	middleNode(head)
}
