package main

func main() {

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list := &ListNode{0, nil}
	result := list
	tmp := 0
	for l1 != nil || l2 != nil || tmp != 0 {
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}

		list.Next = &ListNode{tmp % 10, nil}
		tmp = tmp / 10
		list = list.Next
	}

	return result.Next
}
