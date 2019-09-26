package main

import (
	"fmt"
)

/*
要求：


解题思路：
遍历链表,val小于x时将节点链接上smaller,否则链接上larger,最后将larger链接上smaller,需注意当有larger时要将larger的最后一个节点的next置为nil

关键点：


*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode { // faster 100% less 100%
	if head == nil || head.Next == nil {
		return head
	}
	// smallerHead用于返回,largerHead用于连接上smaller,而smaller和larger始终在迭代
	var smallerHead, smaller, largerHead, larger *ListNode
	if head.Val < x {
		smallerHead = head
		smaller = head
	} else {
		largerHead = head
		larger = head
	}
	head = head.Next
	for head != nil {
		if head.Val < x {
			if smaller == nil {
				smallerHead = head
				smaller = head
			} else {
				smaller.Next = head
				smaller = smaller.Next
			}
		} else {
			if larger == nil {
				largerHead = head
				larger = head
			} else {
				larger.Next = head
				larger = larger.Next
			}
		}
		head = head.Next
	}
	// ListNode中val均小于x时,larger==nil
	if larger != nil {
		larger.Next = nil
	}
	// ListNode中val均大于x时,smaller==nil
	if smaller != nil {
		smaller.Next = largerHead
		return smallerHead
	}
	return largerHead
}

func main() {
	l1 := new(ListNode)
	l1.Val = 2
	l2 := new(ListNode)
	l2.Val = 1
	l1.Next = l2
	// l3 := new(ListNode)
	// l3.Val = 3
	// l2.Next = l3
	// l4 := new(ListNode)
	// l4.Val = 2
	// l3.Next = l4
	// l5 := new(ListNode)
	// l5.Val = 5
	// l4.Next = l5
	// l6 := new(ListNode)
	// l6.Val = 2
	// l5.Next = l6
	result := partition(l1, 2)
	for result != nil {
		fmt.Println(result)
		result = result.Next
	}
}
