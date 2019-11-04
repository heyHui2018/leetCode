package main

import (
	"fmt"
)

/*
要求：
binary search tree
the depth of the two subtrees of every node never differ by more than 1

解题思路：
递归.根据len(nums)/2寻找根节点,则左子树为nums[:mid],右子树为nums[mid+1:],随后以此进行递归即可

关键点：


*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode { // faster 31.91% less 100%
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	t := new(TreeNode)
	t.Val = nums[mid]
	t.Left = sortedArrayToBST(nums[:mid])
	t.Right = sortedArrayToBST(nums[mid+1:])
	return t
}

func main() {
	nums := []int{0, 1, 2, 3, 4, 5}
	result := sortedArrayToBST(nums)
	fmt.Println(result)
}
