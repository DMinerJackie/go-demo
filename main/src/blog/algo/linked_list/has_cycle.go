package main

func main() {

}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast := head.Next
	for fast != nil && head != nil && fast.Next != nil {
		if fast == head {
			return true
		}
		fast = fast.Next.Next
		head = head.Next
	}

	return false
}

func hasCycle1(head *ListNode) bool {
	m := make(map[*ListNode]int)
	for head != nil {
		if _, exist := m[head]; exist {
			return true
		}
		m[head] = 1
		head = head.Next
	}
	return false
}
