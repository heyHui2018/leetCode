package main

import (
	"fmt"
)

/*
要求：


解题思路：
dp[i] = max(dp[i-2]+nums[i],dp[i-1])


关键点：


*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int { // faster 100% less 67.27%

	return nil
}

func main() {
	result := rightSideView(nil)
	fmt.Println(result)
}
