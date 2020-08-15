package main

import "fmt"

/**
反转链表
1、原地反转实现
2、递归实现
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

//反转链表的实现
func reversrList1(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil
	for cur != nil {

		// 注意，这里的pre即表示一个节点，要区别于这里cur实际上是一个链表的表示，因为其有Next，所以pre作为一个节点是可以被指向和指向的
		next := cur.Next // next相当于临时指针，是为了下面cur.Next操作后断了连接，临时保存的关系，方便下一次使用
		cur.Next = pre
		pre = cur
		cur = next

		//pre, cur, cur.Next = cur, cur.Next, pre //这句话最重要
	}
	return pre
}

// 反转链表递归实现1
func reverse(pre, cur *ListNode) *ListNode {
	if cur == nil {
		return pre
	}

	head := reverse(cur, cur.Next)
	cur.Next = pre
	return head
}

func reverseList2(head *ListNode) *ListNode {
	return reverse(nil, head)
}

// 反转链表递归实现2
func reveser3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reveser3(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func main() {
	head := new(ListNode)
	head.Val = 1
	ln2 := new(ListNode)
	ln2.Val = 2
	ln3 := new(ListNode)
	ln3.Val = 3
	ln4 := new(ListNode)
	ln4.Val = 4
	ln5 := new(ListNode)
	ln5.Val = 5
	head.Next = ln2
	ln2.Next = ln3
	ln3.Next = ln4
	ln4.Next = ln5

	//pre := reversrList1(head)
	//for pre != nil {
	//	fmt.Print(pre.Val, " ")
	//	pre = pre.Next
	//}

	//pre2 := reverseList2(head)
	pre2 := reveser3(head)
	for pre2 != nil {
		fmt.Print(pre2.Val, " ")
		pre2 = pre2.Next
	}
}
