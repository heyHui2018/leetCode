package main

import (
	"container/list"
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

// 递归
func generateTrees(n int) []*TreeNode { // faster 100% less 33.33%
	
}

func main() {
	t1 := new(TreeNode)
	t1.Val = 1
	t2 := new(TreeNode)
	t2.Val = 2
	t1.Right = t2
	t3 := new(TreeNode)
	t3.Val = 3
	t2.Left = t3
	result := inorderTraversal(t1)
	fmt.Println(result)
}
