package main

import (
	"fmt"
)

/*
要求：


解题思路：
dp[i]=i
dp[i] = min(dp[i],dp[i-j*j]+1)


关键点：


*/
func numSquares(n int) int { // faster 54.05% less 74.05%
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	result := numSquares(10)
	fmt.Println(result)
}
