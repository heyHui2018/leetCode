package main

import (
	"fmt"
)

/*
要求：


解题思路：
比较当前节点是否为空,值是否相同,再递归比较左子树和右子树

关键点：


*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool { // faster 100% less 50%
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	root := new(TreeNode)
	root.Val = 1
	root.Left = new(TreeNode)
	// root.Right = new(TreeNode)
	root.Left.Val = 2
	// root.Right.Val = 3

	root2 := new(TreeNode)
	root2.Val = 1
	// root2.Left = new(TreeNode)
	root2.Right = new(TreeNode)
	// root2.Left.Val = 2
	root2.Right.Val = 2
	result := isSameTree(root, root2)
	fmt.Println(result)
}
