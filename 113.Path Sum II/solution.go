package main

import (
	"fmt"
)

/*
要求：


解题思路：
递归.递归传递二维数组的指针和一维数组,当节点无左右子节点且值等于剩余sum时,将值加入temp并将temp加入另一数组后再加入结果集,不可将temp直接加入结果集,因为后续还可能对temp进行修改

关键点：


*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int { // faster 98.02% less 75%
	var res [][]int
	r(root, sum, &res, []int{})
	return res
}

func r(root *TreeNode, sum int, res *[][]int, temp []int) {
	if root == nil {
		return
	}
	temp = append(temp, root.Val)
	if root.Left == nil && root.Right == nil && root.Val == sum {
		// 此处必须将temp复制后append进res,否则后续对temp的修改将影响已在res中的值
		cp := append([]int{}, temp...)
		*res = append(*res, cp)
		return
	}
	r(root.Left, sum-root.Val, res, temp)
	r(root.Right, sum-root.Val, res, temp)
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
	sum := -5
	result := pathSum(root, sum)
	for k := range result {
		fmt.Println(result[k])
	}
}
