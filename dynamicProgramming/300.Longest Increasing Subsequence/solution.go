package main

import (
	"fmt"
)

/*
要求：


解题思路：
dp[i]表示以nums[i]结尾的递增子序列的长度，dp中的最大值即为结果
遍历0-i，当nums[i]>nums[j]&&dp[j]+1>dp[i]时，dp[i]=dp[j]+1，即将i更新成其之前最长自增序列的长度


关键点：


*/
func lengthOfLIS(nums []int) int { // faster 64.47% less 34.12%
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	result := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
}

func main() {
	result := lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18, 1})
	fmt.Println(result)
}
