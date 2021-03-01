package main

import (
	"fmt"
)

/*
要求：


解题思路：
后序遍历，左-右-根


关键点：
递归


*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int { // faster 100% less 48.66%
	result := []int{}
	if root == nil {
		return result
	}
	traversal(root, &result)
	return result
}

func traversal(root *TreeNode, list *[]int) {
	if root.Left != nil {
		traversal(root.Left, list)
	}
	if root.Right != nil {
		traversal(root.Right, list)
	}
	*list = append(*list, root.Val)
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
		},
	}
	result := postorderTraversal(root)
	fmt.Println(result)
}
