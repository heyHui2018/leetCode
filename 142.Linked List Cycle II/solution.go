package main

import (
	"fmt"
)

/*
要求：


解题思路：
快慢指针，两指针第一次相遇时，快比慢多跑一圈，故 2 * (head到环起点 + 环起点到相遇点) = head到环起点 + 环起点到相遇点 + 一圈环，
故 head到环起点 + 环起点到相遇点 = 一圈环，故 head到环起点 = 一圈环 - 环起点到相遇点，故 head到环起点 = 相遇点到环起点。
所以此时让慢指针2从起点开始，慢指针1继续，相遇点即为环起点。


关键点：


*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode { // faster 97.53% less 42.59%
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head.Next, head.Next.Next
	for fast != nil && fast.Next != nil && slow != fast {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if slow != fast {
		return nil
	}
	for head != slow {
		head = head.Next
		slow = slow.Next
	}
	return slow
}

func main() {
	head := new(ListNode)
	result := detectCycle(head)
	fmt.Println(result)
}
