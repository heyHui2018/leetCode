package main

import (
	"fmt"
)

/*
要求：


解题思路：
递归.仅有左或右子树时向下递归,均有时分别递归取较小值

关键点：
没有子节点的node才算leaf

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int { // faster 97.01% less 100%
	switch {
	case root == nil:
		return 0
	case root.Left == nil:
		return 1 + minDepth(root.Right)
	case root.Right == nil:
		return 1 + minDepth(root.Left)
	default:
		return 1 + min(minDepth(root.Left), minDepth(root.Right))
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	root := new(TreeNode)
	root.Val = 1
	root.Left = new(TreeNode)
	root.Left.Val = 2
	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 3
	root.Left.Left.Left = new(TreeNode)
	root.Left.Left.Left.Val = 4
	root.Left.Left.Left.Left = new(TreeNode)
	root.Left.Left.Left.Left.Val = 5
	result := minDepth(root)
	fmt.Println(result)
}
