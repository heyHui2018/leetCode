package main

import (
	"fmt"
)

/*
要求：


解题思路：
在102的基础上将其结果的第一维倒叙遍历即可

关键点：


*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int { // faster 51.06% less 100%
	var temp [][]int
	r(&temp, root, 0)
	res := make([][]int, len(temp))
	for k := range res {
		res[k] = temp[len(temp)-1-k]
	}
	return res
}

func r(temp *[][]int, root *TreeNode, level int) {
	if root == nil {
		return
	}
	if len(*temp) <= level {
		*temp = append(*temp, []int{})
	}
	(*temp)[level] = append((*temp)[level], root.Val)
	r(temp, root.Left, level+1)
	r(temp, root.Right, level+1)
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
	result := levelOrderBottom(root)
	for k := range result {
		fmt.Println(result[k])
	}
}
