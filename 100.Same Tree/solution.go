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

func isSameTree(p *TreeNode, q *TreeNode) bool { // faster 60.31% less 100%

}

func main() {
	root := new(TreeNode)
	root.Val = 3
	root.Left = new(TreeNode)
	root.Right = new(TreeNode)
	root.Left.Val = 1
	root.Right.Val = 4
	root.Right.Left = new(TreeNode)
	root.Right.Left.Val = 2
	// root.Right.Left = new(TreeNode)
	// root.Right.Right = new(TreeNode)
	// root.Right.Left.Val = 6
	// root.Right.Right.Val = 20
	result := isSameTree(root, root)
	fmt.Println(result)
}
