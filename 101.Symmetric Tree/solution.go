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

func isSymmetric(root *TreeNode) bool {// faster 100% less 50%

}

func main() {
	root := new(TreeNode)
	root.Val = 1
	root.Left = new(TreeNode)
	// root.Right = new(TreeNode)
	root.Left.Val = 2
	// root.Right.Val = 3
	result := isSymmetric(root)
	fmt.Println(result)
}
