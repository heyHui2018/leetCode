package main

import (
	"fmt"
)

/*
要求：


解题思路：
归并


关键点：


*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode { // faster 100% less 24.02%
	if head == nil || head.Next == nil {
		return head
	}
	left, right := split(head)
	return merge(sortList(left), sortList(right))
}

func split(head *ListNode) (*ListNode, *ListNode) {
	// 快慢指针寻找中点
	slow, fast := head, head
	slowPre := &ListNode{Next: slow}
	for fast != nil && fast.Next != nil {
		slowPre, slow = slow, slow.Next
		fast = fast.Next.Next
	}

	// 截断list
	slowPre.Next = nil
	return head, slow
}

// 合并
func merge(left, right *ListNode) *ListNode {
	cur := &ListNode{}
	headPre := cur
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}
	// 连上left/right剩下的
	if left == nil {
		cur.Next = right
	} else {
		cur.Next = left
	}
	return headPre.Next
}

func main() {
	node := &ListNode{}
	result := sortList(node)
	fmt.Println(result)
}
