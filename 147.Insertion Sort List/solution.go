package main

import (
	"fmt"
)

/*
要求：


解题思路：


关键点：


*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode { // faster 100% less 100%
	headPre := &ListNode{Next: head}

	cur := head
	for cur != nil && cur.Next != nil {
		p := cur.Next
		if p.Val >= cur.Val {
			// p并非插入cur之前
			cur = p
			continue
		}

		// p插入cur之前
		cur.Next = p.Next

		// 从头往后遍历
		pre, next := headPre, headPre.Next

		for next.Val < p.Val {
			pre = next
			next = next.Next
		}

		pre.Next = p
		p.Next = next
	}

	return headPre.Next
}

func main() {
	node := &ListNode{}
	result := insertionSortList(node)
	fmt.Println(result)
}
