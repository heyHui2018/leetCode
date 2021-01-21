package main

import (
	"fmt"
)

/*
要求：


解题思路：
第一次遍历拷贝next，第二次遍历拷贝random

关键点：


*/

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node { // faster 100% less 25%
	if head == nil {
		return nil
	}
	h := head
	nMap := make(map[*Node]*Node)
	for h != nil {
		temp := &Node{Val: h.Val}
		nMap[h] = temp
		h = h.Next
	}
	h = head
	for h != nil {
		nMap[h].Next = nMap[h.Next]
		nMap[h].Random = nMap[h.Random]
		h = h.Next
	}
	return nMap[head]
}

func main() {
	head := new(Node)
	result := copyRandomList(head)
	fmt.Println(result)
}
