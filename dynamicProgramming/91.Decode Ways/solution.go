package main

import (
	"fmt"
)

/*
要求：


解题思路：
取一维数组dp,长度=s+1,dp[i]表示s中前i个字符的解码方式数,取dp[0]=1,当s[1]=0时,因0不对应字符故结果=0,否则dp[1]=1.当s[i]!=0时,dp[i]先赋值dp[i-1],否则赋值0,随后判断s[i-1]是否存在,
若存在且与s[i]组合后>0且<=26,则dp[i]=dp[i-1]+dp[i-2],最后返回dp[len(dp)-1]即可

关键点：
一位数!=0,两位数不大于26且十位!=0

*/

func numDecodings(s string) int { // faster 100% less 100%
	if s[0] == '0' {
		return 0
	}
	n := len(s)
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	lastOne, lastTwo := int(s[0]-'0'), 0
	for i := 2; i <= n; i++ {
		last := int(s[i-1] - '0')
		lastOne, lastTwo = last, lastOne*10+last
		if 1 <= lastOne {
			dp[i] = dp[i-1]
		}
		if 10 <= lastTwo && lastTwo <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}

func main() {
	s := "226"
	result := numDecodings(s)
	fmt.Println(result)
}
