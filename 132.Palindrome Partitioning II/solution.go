package main

import (
	"fmt"
)

/*
要求：


解题思路：
动态规划dp：输出次数而非具体结果

关键点：


*/

func minCut(s string) int { // faster 85.71% less 57.14%
	l := len(s)
	if l == 0 {
		return 0
	}

	isPal := make([][]bool, l) // isPal[i][j]表示字符串s的子串s[i,j]是否为回文串
	for i := 0; i < l; i++ {
		isPal[i] = make([]bool, l)
	}
	cut := make([]int, l) // cut[j]表示子串s[0,j]所需要的最小分割数

	for i := 0; i < l; i++ {
		cut[i] = i // [0,i]最多分割i次
		for j := 0; j <= i; j++ {
			if s[i] == s[j] && (i-j <= 1 || isPal[j+1][i-1]) { // s[j,i]是回文串
				isPal[j][i] = true
				if j == 0 { // 上面初始化了cut[i]需i次，此处需区分s[j,i]是否是回文以将cut[i]置0（因s[j,i-1]非回文时s[j,i]可能是回文，故cut[i]不能简单的等于cut[i-1]+1）
					cut[i] = 0
				} else {
					cut[i] = min(cut[i], cut[j-1]+1)
				}
			}
		}
	}
	return cut[l-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	s := "aab"
	result := minCut(s)
	fmt.Println(result)
}
