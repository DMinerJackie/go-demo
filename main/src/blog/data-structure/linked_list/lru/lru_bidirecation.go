package lru

//type LRUNode struct {
//	Key   int
//	Value int
//	Pre   *LRUNode
//	Next  *LRUNode
//}
//
//type LRU struct {
//	capacity int
//	dataMap  map[int]*LRUNode
//	head     *LRUNode
//	tail     *LRUNode
//	count    int
//}
//
//// 初始化
//func Construct(capacity int) LRU {
//	head := &LRUNode{}
//	tail := &LRUNode{}
//	head.Next = tail
//	tail.Pre = head
//	return LRU{capacity: capacity, count: 0, dataMap: make(map[int]*LRUNode), head: head, tail: tail}
//}
//
//// 插入节点
//func (lru *LRU) insertFront(node *LRUNode) {
//	node.Next = lru.head.Next
//	lru.head.Next.Pre = node
//	lru.head.Next = node
//	node.Pre = lru.head
//}
//
//// 删除节点
//func (lru *LRU) detachNode(node *LRUNode) {
//	node.Pre.Next = node.Next
//	node.Next.Pre = node.Pre
//}
//
//func (lru *LRU) delLast() {
//
//}
//
//func (lru *LRU) addNode(node *LRUNode) {
//	lru.dataMap[node.Key] = node
//	lru.head.Next = node
//	lru.head = lru.head.Next
//
//}

// https://leetcode-cn.com/problems/lru-cache/solution/golang-lru-cache-dai-ma-chao-ji-rong-yi-li-jie-by-/

type LRUCache struct {
	Cap  int
	Map  map[int]*Node
	Head *Node
	Last *Node
}

type Node struct {
	Val  int
	Key  int
	Pre  *Node
	Next *Node
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		Cap:  capacity,
		Map:  make(map[int]*Node, capacity),
		Head: &Node{},
		Last: &Node{},
	}
	cache.Head.Next = cache.Last
	cache.Last.Pre = cache.Head
	return cache
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.Map[key]
	if !ok {
		return -1
	}
	this.remove(node)
	this.setHeader(node)
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.Map[key]
	if ok {
		this.remove(node)
	} else {
		if len(this.Map) == this.Cap {
			delete(this.Map, this.Last.Pre.Key)
			this.remove(this.Last.Pre)
		}
		node = &Node{Val: value, Key: key}
		this.Map[node.Key] = node
	}
	node.Val = value
	this.setHeader(node)
}

func (this *LRUCache) setHeader(node *Node) {
	this.Head.Next.Pre = node
	node.Next = this.Head.Next
	this.Head.Next = node
	node.Pre = this.Head
}

func (this *LRUCache) remove(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func main() {

}
