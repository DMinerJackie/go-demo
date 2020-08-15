package main

func main() {

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	result := &ListNode{}
	result.Next = head
	var pre *ListNode
	cur := result
	i := 1

	for head != nil {
		if i >= n {
			pre = cur
			cur = cur.Next
		}

		head = head.Next
		i++
	}

	pre.Next = pre.Next.Next
	return result.Next
}
