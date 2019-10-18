package main

import (
	"fmt"
)

/*
要求：


解题思路：
动态规划.dp[i][j] == true 表示 s1[:i] 和 s2[:j] 可以生成 s3[:i+j]

关键点：


*/

func isInterleave(s1 string, s2 string, s3 string) bool { // faster 100% less 50%
	if len(s3) != (len(s1) + len(s2)) {
		return false
	}
	dp := make([][]bool, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		dp[i] = make([]bool, len(s2)+1)
	}
	dp[0][0] = true
	for i := 1; i <= len(s1); i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}
	for j := 1; j <= len(s2); j++ {
		dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
	}
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}
	return dp[len(s1)][len(s2)]
}

func main() {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"
	result := isInterleave(s1, s2, s3)
	fmt.Println(result)
}
