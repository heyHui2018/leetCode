package main

import (
	"fmt"
)

/*
要求：


解题思路：
递归无法满足时间要求
动态规划,dp[i][j]表示s[:j]能组成t[:i]的子序列个数,当t[i]==s[j]时,dp[i][j]=dp[i][j-1]+dp[i-1][j-1],当t[i-1]!=s[j-1]时,dp[i][j]=dp[i][j-1]

关键点：


*/

// 递归
func numDistinct1(s string, t string) int { // Time Limit Exceeded
	if len(s) < len(t) {
		return 0
	}
	var res [][]int
	r1(s, t, 0, 0, &res, []int{})
	for k := range res {
		fmt.Println(res[k])
	}
	return len(res)
}

func r1(s, t string, sIndex, tIndex int, res *[][]int, temp []int) {
	for ; sIndex < len(s); sIndex++ {
		if s[sIndex] == t[tIndex] {
			temp = append(temp, sIndex)
			if tIndex == len(t)-1 {
				tt := append([]int{}, temp...)
				*res = append(*res, tt)
				// 组成t后去掉最后一位继续在s上往后遍历
				r1(s, t, sIndex+1, tIndex, res, temp[:len(temp)-1])
				return
			}
			r1(s, t, sIndex+1, tIndex+1, res, temp)
			// 回溯
			temp = temp[:len(temp)-1]
		}
	}
}

// 动态规划
func numDistinct(s string, t string) int { // faster 100% less 50%
	// 第一维t,第二维s
	dp := make([][]int, len(t)+1)
	for k := range dp {
		dp[k] = make([]int, len(s)+1)
	}
	// t为空时,无论s如何,结果均为0
	for k := 0; k < len(s); k++ {
		dp[0][k] = 1
	}

	for i := 1; i < len(t)+1; i++ {
		for j := 1; j < len(s)+1; j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i][j-1] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[len(t)][len(s)]
}

func main() {
	s := "rabbbit"
	t := "rabbit"
	result := numDistinct(s, t)
	fmt.Println(result)
}
