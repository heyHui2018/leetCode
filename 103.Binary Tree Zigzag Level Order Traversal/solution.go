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

func zigzagLevelOrder(root *TreeNode) [][]int { // faster 36.36% less 50%

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
	result := levelOrder(root)
	for k := range result {
		fmt.Println(result[k])
	}
}
