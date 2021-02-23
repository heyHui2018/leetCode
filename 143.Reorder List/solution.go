package main

import (
	"fmt"
	"time"
)

/*
要求：


解题思路：
1、找到中点；
2、翻转后半段；
3、合并。

关键点：


*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) { // faster 96.75% less 42.21%
	if head == nil || head.Next == nil {
		return
	}

	// 寻找中点
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	next := slow.Next // slow.Next为后半段起点

	// 翻转后半段
	for next != nil {
		temp := next.Next
		next.Next = slow
		slow = next
		next = temp
	}
	end := slow // end为原链表最后节点

	// 合并
	for head != end {
		hNext := head.Next
		eNext := end.Next
		head.Next = end  // head第一节点指向end第一节点
		end.Next = hNext // end第一节点指向head第二节点
		head = hNext     // head迭代
		end = eNext      // end迭代
	}
	end.Next = nil
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}
	reorderList(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
		time.Sleep(100 * time.Millisecond)
	}
}
