package main

import "fmt"

type ListNode1 struct {
	data interface{}
	Next *ListNode1
}

//反转单链表
func reverseList(head *ListNode1) *ListNode1 {
	cur := head
	var pre *ListNode1
	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}
	return pre
}

func CreateNode(node *ListNode1, max int) {
	cur := node
	for i := 1; i < max; i++ {
		cur.Next = &ListNode1{}
		cur.Next.data = i
		cur = cur.Next
	}
}

//打印链表的方法
func PrintNode(info string, node *ListNode1) {
	fmt.Print(info)
	for cur := node; cur != nil; cur = cur.Next {
		fmt.Print(cur.data, " ")
	}
	fmt.Println()
}

type student struct {
	Name string
	Age  int
}

func main() {

	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	var head = new(ListNode1)
	CreateNode(head, 10)
	PrintNode("前：", head)
	yyy := reverseList(head)
	PrintNode("后：", yyy)

}
