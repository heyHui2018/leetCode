package main

import (
	"fmt"
)

/*
要求：
非递归

解题思路：


关键点：


*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int { // faster 5.95% less 100%

}

func main() {
	root := new(TreeNode)
	result := inorderTraversal(root)
	fmt.Println(result)
}
