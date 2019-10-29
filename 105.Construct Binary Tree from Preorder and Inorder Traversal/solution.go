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

func buildTree(preorder []int, inorder []int) *TreeNode { // faster 94.84% less 83.33%

}

func main() {
	preorder := []int{}
	inorder := []int{}
	result := buildTree(preorder, inorder)
	fmt.Println(result)
}
