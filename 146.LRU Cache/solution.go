package main

/*
要求：


解题思路：
可通过golang自带的container/list实现


关键点：


*/

type LRUCache struct {
	Capacity int
	Cache    map[int]*LRUCacheNode
	head     *LRUCacheNode
	tail     *LRUCacheNode
}

type LRUCacheNode struct {
	Key   int
	Value int
	Prev  *LRUCacheNode
	Next  *LRUCacheNode
}

func Constructor(capacity int) LRUCache { // faster 96.83% less 60.83%
	lru := LRUCache{
		Capacity: capacity,
		Cache:    map[int]*LRUCacheNode{},
		head:     &LRUCacheNode{},
		tail:     &LRUCacheNode{},
	}
	lru.head.Next = lru.tail
	lru.tail.Prev = lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Cache[key]; ok {
		this.move2Head(node)
		return node.Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Cache[key]; ok {
		node.Value = value
		this.move2Head(node)
		return
	}

	node := &LRUCacheNode{
		Key:   key,
		Value: value,
	}
	this.Cache[key] = node
	this.add2Head(node)
	if len(this.Cache) > this.Capacity {
		node2Remove := this.removeTail()
		delete(this.Cache, node2Remove.Key)
	}
}

func (this *LRUCache) move2Head(node *LRUCacheNode) {
	this.removeNode(node)
	this.add2Head(node)
}

func (this *LRUCache) removeNode(node *LRUCacheNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) add2Head(node *LRUCacheNode) {
	node.Prev = this.head
	node.Next = this.head.Next
	this.head.Next.Prev = node
	this.head.Next = node
}

func (this *LRUCache) removeTail() *LRUCacheNode {
	node := this.tail.Prev
	this.removeNode(node)
	return node
}
