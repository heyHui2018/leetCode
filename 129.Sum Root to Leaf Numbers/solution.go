package main

import (
	"fmt"
)

/*
要求：


解题思路：
先将array转为map，随后遍历map，判断是否存在当前数字的前一位，若存在则继续遍历，否则开始往后寻找所有连续的序列。
此法仅需处理各个序列的起点。

关键点：


*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int { // faster 98。35% less 63.7%

	return 0
}

func main() {
	root := &TreeNode{

	}
	result := sumNumbers(root)
	fmt.Println(result)
}
