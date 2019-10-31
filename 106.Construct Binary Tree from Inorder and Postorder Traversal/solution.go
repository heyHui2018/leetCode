package main

import (
	"fmt"
)

/*
要求：


解题思路：
递归.postorder[len-1]为根节点,找到inorder中postorder[len-1]的位置i,postorder[:i]和inorder[:i]为左子树,postorder[i:len(postorder)-1]和inorder[i+1:]为右子树,随后继续递归

关键点：


*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode { // faster 95.65% less 25%
	if len(inorder) == 0 {
		return nil
	}
	t := new(TreeNode)
	t.Val = postorder[len(postorder)-1]
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
	t.Left = buildTree(inorder[:i], postorder[:i])
	t.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])
	return t
}

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	result := buildTree(inorder, postorder)
	fmt.Println(result)
}
