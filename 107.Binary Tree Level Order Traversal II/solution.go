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

func levelOrderBottom(root *TreeNode) [][]int { // faster 95.65% less 25%


}

func main() {
	root := new(TreeNode)
	result := levelOrderBottom(root)
	for k := range result {
		fmt.Println(result[k])
	}
}