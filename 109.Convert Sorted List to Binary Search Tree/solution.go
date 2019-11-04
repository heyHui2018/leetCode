package main

import (
	"fmt"
)

/*
要求：


解题思路：
递归.通过a,b:=begin,begin;a=a.Next.Next,b=b.Next来快速迭代获得链表的中点即树的根,则左子树为begin到b,右子树为b.Next到end

关键点：


*/

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode { // faster 68.67% less 100%
	return r(head, nil)
}

func r(begin, end *ListNode) *TreeNode {
	if begin == end {
		return nil
	}
	if begin.Next == end {
		return &TreeNode{Val: begin.Val}
	}
	fast, slow := begin, begin
	for fast != end && fast.Next != end {
		fast = fast.Next.Next
		slow = slow.Next
	}
	mid := slow
	t := new(TreeNode)
	t.Val = mid.Val
	t.Left = r(begin, mid)
	t.Right = r(mid.Next, end)
	return t
}

func main() {
	head := new(ListNode)
	result := sortedListToBST(head)
	fmt.Println(result)
}
