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

func deleteDuplicates(head *ListNode) *ListNode { // faster 100% less 33.33%

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
	l4 := new(ListNode)
	l4.Val = 3
	l3.Next = l4
	l5 := new(ListNode)
	l5.Val = 3
	l4.Next = l5
	l6 := new(ListNode)
	l6.Val = 4
	l5.Next = l6
	l7 := new(ListNode)
	l7.Val = 4
	l6.Next = l7
	l8 := new(ListNode)
	l8.Val = 5
	l7.Next = l8
	result := deleteDuplicates(l1)
	for result.Next != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
