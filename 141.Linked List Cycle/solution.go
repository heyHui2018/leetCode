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

func hasCycle(head *ListNode) bool { // faster 100% less 8.56%
	if head == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil && slow != fast {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow == fast
}

func main() {
	head := new(ListNode)
	result := hasCycle(head)
	fmt.Println(result)
}
