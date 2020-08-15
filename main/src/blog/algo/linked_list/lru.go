package main

type LinkNode struct {
	key, value int
	pre, next  *LinkNode
}

type LRUCache struct {
	m          map[int]*LinkNode
	cap        int
	head, tail *LinkNode
}

func Constructor(capacity int) LRUCache {
	head := &LinkNode{0, 0, nil, nil}
	tail := &LinkNode{0, 0, nil, nil}
	head.next = tail
	tail.pre = head
	return LRUCache{make(map[int]*LinkNode), capacity, head, tail}
}

func (this *LRUCache) moveToHead(node *LinkNode) {
	head := this.head
	// 从当前位置删除
	node.pre.next = node.next
	node.next.pre = node.pre
	// 移动到首位置
	node.next = head.next
	head.next.pre = node
	node.pre = head
	head.next = node
}

func (this *LRUCache) Get(key int) int {
	cache := this.m
	if v, exist := cache[key]; exist {
		this.moveToHead(v)
		return v.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	head := this.head
	tail := this.tail
	cache := this.m

	if v, exist := cache[key]; exist {
		v.value = value
		this.moveToHead(v)
	} else {
		v := &LinkNode{key, value, nil, nil}
		if len(cache) == this.cap {
			delete(cache, tail.pre.key)
			tail.pre.pre.next = tail
			tail.pre = tail.pre.pre
		}

		v.next = head.next
		v.pre = head
		head.next.pre = v
		head.next = v
		cache[key] = v
	}

}
