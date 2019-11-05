package main

import (
	"fmt"
)

/*
要求：
in-place

解题思路：


关键点：


*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) { // faster 98.02% less 75%

}

func main() {
	root := new(TreeNode)
	root.Val = -2
	root.Left = new(TreeNode)
	root.Left.Val = -3
	root.Right = new(TreeNode)
	root.Right.Val = -3
	// root.Left.Left = new(TreeNode)
	// root.Left.Left.Val = 3
	// root.Left.Left.Left = new(TreeNode)
	// root.Left.Left.Left.Val = 4
	// root.Left.Left.Left.Left = new(TreeNode)
	// root.Left.Left.Left.Left.Val = 5
	flatten(root)
	fmt.Println(root)
}
