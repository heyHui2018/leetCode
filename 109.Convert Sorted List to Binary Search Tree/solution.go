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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode { // faster 31.91% less 100%

}

func main() {
	head := new(ListNode)
	result := sortedListToBST(head)
	fmt.Println(result)
}
