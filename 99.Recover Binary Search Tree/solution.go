package main

import (
	"fmt"
)

/*
要求：


解题思路：
1、中序遍历后将节点放入数组,随后将数组排序,最后将数组转换成树
2、在1的基础上改进,中序遍历,因有两节点位置错误,故遍历中能遇到1或2次降序,若仅一次,则直接交换两节点,若两次,则交换第一次的较大点与第二次的较小点

关键点：


*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) { // faster 60.31% less 100%
	var first, second, prev *TreeNode
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if prev != nil && prev.Val > root.Val {
			if first == nil {
				first = prev
			}
			if first != nil {
				second = root
			}
		}
		prev = root
		dfs(root.Right)
	}
	dfs(root)
	first.Val, second.Val = second.Val, first.Val
}

func main() {
	root := new(TreeNode)
	root.Val = 3
	root.Left = new(TreeNode)
	root.Right = new(TreeNode)
	root.Left.Val = 1
	root.Right.Val = 4
	root.Right.Left = new(TreeNode)
	root.Right.Left.Val = 2
	// root.Right.Left = new(TreeNode)
	// root.Right.Right = new(TreeNode)
	// root.Right.Left.Val = 6
	// root.Right.Right.Val = 20
	recoverTree(root)
	fmt.Println(root, root.Left, root.Right)
}
