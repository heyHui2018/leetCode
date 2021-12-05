package main

import (
	"fmt"
)

/*
要求：


解题思路：
dp[i][j]表示第一个字符串前i个字符到第二个字符串前j个字符需要的编辑距离;若s1[i]==s2[j],dp[i][j]=min(修改:dp[i-1][j-1],删除:dp[i-1][j]+1,增加:dp[i][j-1]+1),
若s1[i]!=s2[j],dp[i][j]=min(修改:dp[i-1][j-1],删除:dp[i-1][j],增加:dp[i][j-1])+1,dp[i][0]=i;i=1...m.dp[0][j]=j;j=1...n

关键点：


*/

func minDistance(word1 string, word2 string) int { // faster 77.99% less 100%
	m := len(word1)
	n := len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1])
			replace := 1
			if word1[i-1] == word2[j-1] {
				replace = 0
			}
			dp[i][j] = min(dp[i][j], dp[i-1][j-1]+replace)
		}
	}
	return dp[m][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	word1 := "intention"
	word2 := "execution"
	result := minDistance(word1, word2)
	fmt.Println(result)
}
