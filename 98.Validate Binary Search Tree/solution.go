package main

import (
	"fmt"
)

/*
要求：


解题思路：
递归,同时传入当前根节点应属于的大小范围,左子树及其子树应小于根节点,右子树同理,递归时不断更新大小范围

关键点：
递归时传入当前根节点应属于的大小范围

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool { // faster 100% less 50%
	MIN, MAX := -1<<63, 1<<63-1
	return r(MIN, MAX, root)
}

func r(min, max int, root *TreeNode) bool {
	if root == nil {
		return true
	}
	return min < root.Val && root.Val < max && r(min, root.Val, root.Left) && r(root.Val, max, root.Right)
}
func main() {
	root := new(TreeNode)
	root.Val = 10
	root.Left = new(TreeNode)
	root.Right = new(TreeNode)
	root.Left.Val = 5
	root.Right.Val = 15
	root.Right.Left = new(TreeNode)
	root.Right.Right = new(TreeNode)
	root.Right.Left.Val = 6
	root.Right.Right.Val = 20
	result := isValidBST(root)
	fmt.Println(result)
}
