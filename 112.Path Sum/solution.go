package main

import (
	"fmt"
)

/*
要求：


解题思路：
递归.当没有左右子节点且root.Val==sum时,返回true;当root==nil时返回false;最终返回左右子节点的或运算值

关键点：
root.Val可能为负数,故无需比较root.Val与sum的大小来提前终止此次探索

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool { // faster 96.63% less 100%
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && root.Val == sum {
		return true
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

func main() {
	root := new(TreeNode)
	root.Val = -2
	root.Left = new(TreeNode)
	root.Left.Val = -3
	// root.Left.Left = new(TreeNode)
	// root.Left.Left.Val = 3
	// root.Left.Left.Left = new(TreeNode)
	// root.Left.Left.Left.Val = 4
	// root.Left.Left.Left.Left = new(TreeNode)
	// root.Left.Left.Left.Left.Val = 5
	sum := -5
	result := hasPathSum(root, sum)
	fmt.Println(result)
}
