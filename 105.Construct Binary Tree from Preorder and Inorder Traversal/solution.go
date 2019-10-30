package main

import (
	"fmt"
)

/*
要求：


解题思路：
递归.preorder[0]为根节点,找到inorder中preorder[0]的位置i,preorder[1:i+1]和inorder[:i]为左子树,preorder[i+1:]和inorder[i+1:]为右子树,随后继续递归

关键点：


*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode { // faster 50.80% less 83.33%
	if len(inorder) == 0 {
		return nil
	}
	t := new(TreeNode)
	t.Val = preorder[0]
	if len(inorder) == 1 {
		return t
	}
	i := 0
	for k := range inorder {
		if inorder[k] == t.Val {
			i = k
			break
		}
	}
	t.Left = buildTree(preorder[1:i+1], inorder[:i])
	t.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return t
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	result := buildTree(preorder, inorder)
	fmt.Println(result)
}
