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

func buildTree(inorder []int, postorder []int) *TreeNode { // faster 50.80% less 83.33%

}

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	result := buildTree(inorder, postorder)
	fmt.Println(result)
}
