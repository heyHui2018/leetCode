package main

import (
	"fmt"
)

/*
要求：


解题思路：
同时遍历l1和l2,将对应值相加,有进位时放入carry,同时对10取余获得对应l3的值,最终输出l3

关键点：
(2 -> 4 -> 3)是342

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resPre := &ListNode{}
	cur := resPre
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10

		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}

	return resPre.Next
}

func main() {
	l1 := new(ListNode)
	l1.Val = 5
	l1.Next = new(ListNode)
	l1.Next.Val = 4
	// l1.Next.Next = new(ListNode)
	// l1.Next.Next.Val = 3

	l2 := new(ListNode)
	l2.Val = 5
	l2.Next = new(ListNode)
	l2.Next.Val = 9
	l2.Next.Next = new(ListNode)
	l2.Next.Next.Val = 4

	result := addTwoNumbers(l1, l2)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
