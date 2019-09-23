package main

import (
	"fmt"
	"time"
)

/*
要求：


解题思路：
遍历链表,取tempVal记录临时val,tempNode记录当前最后一个有效节点,当发现当前val=tempVal时,若当前节点next==nil,则将tempNode的next置为nil以使其成为最后一个节点,若!=nil,则head = head.Next继续遍历;
当val!=tempVal时,更新tempVal,将当前节点链接上tempNode,并更新tempNode,后续处理同上

关键点：


*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode { // faster 89.32% less 100%
	if head == nil || head.Next == nil {
		return head
	}
	res := head
	tempVal := head.Val
	tempNode := head
	head = head.Next
	for head != nil {
		if head.Val == tempVal {
			if head.Next == nil {
				tempNode.Next = nil
				return res
			}
			head = head.Next
			continue
		}
		tempVal = head.Val
		tempNode.Next = head
		tempNode = head
		if head.Next == nil {
			tempNode.Next = nil
			return res
		}
		head = head.Next
	}
	return res
}

func main() {
	l1 := new(ListNode)
	l1.Val = 1
	l2 := new(ListNode)
	l2.Val = 1
	l1.Next = l2
	l3 := new(ListNode)
	l3.Val = 2
	l2.Next = l3
	// l4 := new(ListNode)
	// l4.Val = 3
	// l3.Next = l4
	// l5 := new(ListNode)
	// l5.Val = 3
	// l4.Next = l5
	// l6 := new(ListNode)
	// l6.Val = 4
	// l5.Next = l6
	// l7 := new(ListNode)
	// l7.Val = 4
	// l6.Next = l7
	// l8 := new(ListNode)
	// l8.Val = 5
	// l7.Next = l8
	// l9 := new(ListNode)
	// l9.Val = 5
	// l8.Next = l9
	result := deleteDuplicates(l1)
	for result.Next != nil {
		time.Sleep(1 * time.Second)
		fmt.Println(result)
		result = result.Next
	}
	fmt.Println(result)
}
