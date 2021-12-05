package main

import (
	"fmt"
)

/*
要求：


解题思路：
f(n)的根的情况有n种(1~n),对于根i,左子树取值1~i-1(即f(i-1)),右子树取值i+1~n(即f(n-i),但各值需+i),则f(n) = ∑f(i-1)*f(n-i) (i=1...n)

关键点：
binary search trees 二分搜索树 左子树值<根节点值<右子树值

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode { // faster 61.04% less 100%
	if n == 0 {
		return []*TreeNode{}
	}
	return gTrees(1, n)
}

func gTrees(left, right int) []*TreeNode {
	if left > right {
		return []*TreeNode{nil}
	} else if left == right {
		// 相等时表示仅有根节点
		return []*TreeNode{&TreeNode{Val: left}}
	}
	var res []*TreeNode
	// 根节点的n种可能
	for num := left; num <= right; num++ {
		// 左子树所有可能
		leftTrees := gTrees(left, num-1)
		// 右子树所有可能
		rightTrees := gTrees(num+1, right)
		// 整合
		for i := range leftTrees {
			for j := range rightTrees {
				root := &TreeNode{Val: num}
				root.Left = leftTrees[i]
				root.Right = rightTrees[j]
				res = append(res, root)
			}
		}
	}
	return res
}

func main() {
	n := 3
	result := generateTrees(n)
	for k := range result {
		fmt.Println(result[k])
	}
}
