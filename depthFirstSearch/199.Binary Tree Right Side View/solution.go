package main

import (
	"fmt"
)

/*
要求：


解题思路：
二叉树的右视图，定义a、b、c三个数组，a存放当前层的节点，b存放a中节点的子字节，c存放结果
取b中最右节点的值放入c，随后将b复制进a，清空b


关键点：


*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int { // faster 100% less 20.35%
	if root == nil {
		return []int{}
	}
	var a []*TreeNode
	a = append(a, root)
	result := []int{root.Val}
	for len(a) > 0 {
		var b []*TreeNode
		for i := range a {
			if a[i].Left != nil {
				b = append(b, a[i].Left)
			}
			if a[i].Right != nil {
				b = append(b, a[i].Right)
			}
		}
		if len(b) == 0 {
			return result
		}
		result = append(result, b[len(b)-1].Val)

		a = make([]*TreeNode, len(b))
		for i := range b {
			a[i] = b[i]
		}
		b = []*TreeNode{}
	}
	return result
}

func main() {
	// [1,2,3,null,5,null,4]
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	result := rightSideView(root)
	fmt.Println(result)
}
