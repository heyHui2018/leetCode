package main

import (
	"fmt"
)

/*
要求：


解题思路：
递归.递归时传入切片指针及深度(即所属二维切片res中的第几个切片),当len(res)<level时,先初始化属于当前level的第二维切片,当level%2==0即偶数时,照常添加,奇数时,将res中对应level第二维切片内的数组添加到当前root的val后

关键点：


*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int { // faster 100% less 100%
	var res [][]int
	r(0, root, &res)
	return res
}

func r(level int, root *TreeNode, res *[][]int) {
	if root == nil {
		return
	}
	if len(*res) <= level {
		for i := len(*res); i <= level; i++ {
			*res = append(*res, []int{})
		}
	}
	if level%2 == 0 {
		(*res)[level] = append((*res)[level], root.Val)
	} else {
		temp := make([]int ,  len((*res)[level])+1)
		temp[0] = root.Val
		copy(temp[1:], (*res)[level])
		(*res)[level] = temp
	}
	r(level+1, root.Left, res)
	r(level+1, root.Right, res)
	return
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
	result := zigzagLevelOrder(root)
	for k := range result {
		fmt.Println(result[k])
	}
}
