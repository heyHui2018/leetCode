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

func hasPathSum(root *TreeNode, sum int) bool { // faster 97.01% less 100%

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
	sum := 5
	result := hasPathSum(root, sum)
	fmt.Println(result)
}
