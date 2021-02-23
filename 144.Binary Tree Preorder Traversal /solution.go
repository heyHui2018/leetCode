package main

import (
	"fmt"
)

/*
要求：


解题思路：


关键点：


*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int { // faster 100% less 39。35%
	result := []int{}
	if root == nil {
		return result
	}
	traversal(root, &result)
	return result
}

func traversal(root *TreeNode, list *[]int) {
	*list = append(*list, root.Val)
	if root.Left != nil {
		traversal(root.Left, list)
	}
	if root.Right != nil {
		traversal(root.Right, list)
	}
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
	result := preorderTraversal(root)
	fmt.Println(result)
}
