package main

import (
	"fmt"
)

/*
要求：


解题思路：
打劫第一个和不打劫，两种结果进行比较
dp[i] = max(dp[i-2]+nums[i],dp[i-1])


关键点：


*/

func rob(nums []int) int { // faster 100% less 67.16%
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	// 打劫第一家
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums)-1; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	a := dp[len(nums)-2]

	// 不打劫第一家
	dp = make([]int, len(nums))
	dp[1] = nums[1]
	dp[2] = max(nums[1], nums[2])
	for i := 3; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	b := dp[len(nums)-1]
	return max(a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	result := rob([]int{1, 2, 3, 1})
	fmt.Println(result)
}
