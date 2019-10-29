package main

import (
	"fmt"
)

/*
要求：


解题思路：
递归.root == nil时返回,取两个子节点中值较大者返回

关键点：


*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int { // faster 94.84% less 83.33%
	return r(root, 0)
}

func r(root *TreeNode, level int) int {
	if root == nil {
		return level
	}
	return max(r(root.Left, level+1), r(root.Right, level+1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := new(TreeNode)
	root.Val = 3
	root.Left = new(TreeNode)
	root.Right = new(TreeNode)
	root.Left.Val = 9
	root.Right.Val = 20
	root.Right.Left = new(TreeNode)
	root.Right.Right = new(TreeNode)
	root.Right.Left.Val = 15
	root.Right.Right.Val = 7
	root.Right.Right.Right = new(TreeNode)
	root.Right.Right.Right.Val = 1
	result := maxDepth(root)
	fmt.Println(result)
}
