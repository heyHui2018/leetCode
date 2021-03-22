package main

import (
	"fmt"
)

/*
要求：


解题思路：
分别遍历A、B获取长度，计算长度差后使节点较多的链表去掉其头部节点，随后同时往后遍历

关键点：


*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode { // faster 86.05% less 26.16%
	temp := headA
	lenA := 0
	for temp != nil {
		lenA++
		temp = temp.Next
	}
	temp = headB
	lenB := 0
	for temp != nil {
		lenB++
		temp = temp.Next
	}

	diff := lenA - lenB
	switch {
	case diff > 0:
		for diff > 0 {
			headA = headA.Next
			diff--
		}
	case diff < 0:
		for diff < 0 {
			headB = headB.Next
			diff++
		}
	}

	for headA != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}

func main() {
	same := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	}
	headA := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: same,
		},
	}
	headB := &ListNode{
		Val:  2,
		Next: same,
	}
	result := getIntersectionNode(headA, headB)
	fmt.Println(result)
}
