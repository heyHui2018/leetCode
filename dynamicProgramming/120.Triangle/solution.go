package main

import (
	"fmt"
)

/*
要求：


解题思路：
动态规划，第n行的第x个数，相邻的是第n+1行的第x和第x+1个数，min(dp[n][x]+dp[n+1][x],dp[n][x]+dp[n+1][x+1]),
所以二维的dp可优化成一维，min(dp[x],dp[x+1])，由上往下的最短路径，需从下往上迭代，初始化赋值最后一行，随后从倒数第二行开始迭代

关键点：


*/

func minimumTotal(triangle [][]int) int { // faster 93.98% less 100%
	if len(triangle) == 0 {
		return 0
	}
	// 初始化,赋值triangle最后一行
	dp := make([]int, len(triangle))
	for k := range dp {
		dp[k] = triangle[len(triangle)-1][k]
	}
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := range triangle[i] {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}
	return dp[0]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	triangle := [][]int{
		[]int{2}, []int{3, 4}, []int{6, 7, 5}, []int{4, 1, 8, 3},
	}
	result := minimumTotal(triangle)
	fmt.Println(result)
}
