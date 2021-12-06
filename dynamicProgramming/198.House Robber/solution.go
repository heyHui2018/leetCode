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

func rob(nums []int) int { // faster 100% less 67.27%
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{-1, -2}
	result := rob(nums)
	fmt.Println(result)
}
