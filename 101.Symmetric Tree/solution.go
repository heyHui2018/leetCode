package main

import (
	"fmt"
)

/*
要求：


解题思路：
与100Same Tree类似,不同点在于从根节点往下分成两棵树后,要判断的是左子树的左子树与右子树的右子树是否same,左子树的右子树与右子树的左子树是否same

关键点：


*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool { // faster 100% less 50%
	if root == nil {
		return true
	}
	return r(root.Left, root.Right)
}

func r(p, q *TreeNode) bool {
	fmt.Println(p, q)
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && r(p.Left, q.Right) && r(p.Right, q.Left)
}

func main() {
	root := new(TreeNode)
	root.Val = 1
	root.Left = new(TreeNode)
	root.Right = new(TreeNode)
	root.Left.Val = 2
	root.Right.Val = 2
	root.Left.Left = new(TreeNode)
	root.Left.Right = new(TreeNode)
	root.Left.Left.Val = 3
	root.Left.Right.Val = 4
	root.Right.Left = new(TreeNode)
	root.Right.Right = new(TreeNode)
	root.Right.Left.Val = 4
	root.Right.Right.Val = 3
	root.Left.Left.Left = new(TreeNode)
	root.Left.Left.Right = new(TreeNode)
	root.Left.Left.Left.Val = 5
	root.Left.Left.Right.Val = 6
	root.Left.Right.Left = new(TreeNode)
	root.Left.Right.Right = new(TreeNode)
	root.Left.Right.Left.Val = 7
	root.Left.Right.Right.Val = 8
	root.Right.Left.Left = new(TreeNode)
	root.Right.Left.Right = new(TreeNode)
	root.Right.Left.Left.Val = 8
	root.Right.Left.Right.Val = 7
	root.Right.Right.Left = new(TreeNode)
	root.Right.Right.Right = new(TreeNode)
	root.Right.Right.Left.Val = 6
	root.Right.Right.Right.Val = 5
	result := isSymmetric(root)
	fmt.Println(result)
}
