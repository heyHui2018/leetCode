package main

import (
	"fmt"
)

/*
要求：


解题思路：
递归.当左右子树均平衡且左右子树深度差不大于1时,向上返回true和最大深度+1,否则返回false

关键点：


*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool { // faster 99.37% less 50%
	_, isBalanced := r(root)
	return isBalanced
}

func r(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	leftDepth, leftIsBalanced := r(root.Left)
	rightDepth, rightIsBalanced := r(root.Right)

	if leftIsBalanced && rightIsBalanced &&
		-1 <= leftDepth-rightDepth && leftDepth-rightDepth <= 1 {
		return max(leftDepth, rightDepth) + 1, true
	}

	return 0, false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := new(TreeNode)
	root.Val = 1
	root.Left = new(TreeNode)
	root.Left.Val = 2
	root.Right = new(TreeNode)
	root.Right.Val = 2
	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 3
	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 3
	root.Left.Left.Left = new(TreeNode)
	root.Left.Left.Left.Val = 4
	root.Left.Left.Right = new(TreeNode)
	root.Left.Left.Right.Val = 4
	result := isBalanced(root)
	fmt.Println(result)
}
