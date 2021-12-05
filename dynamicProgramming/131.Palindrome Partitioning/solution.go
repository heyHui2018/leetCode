package main

import (
	"fmt"
)

/*
要求：


解题思路：
广度优先遍历dfs：输出所有可能结果

关键点：


*/

func partition(s string) [][]string { // faster 84.73% less 82.58%
	var result [][]string
	cur := make([]string, 0, len(s))
	dfs(s, 0, cur, &result)
	return result
}

func dfs(s string, i int, cur []string, result *[][]string) {
	if i == len(s) {
		tmp := make([]string, len(cur))
		copy(tmp, cur)
		*result = append(*result, tmp)
		return
	}

	for j := i; j < len(s); j++ {
		// 当s[i : j+1]是回文时才继续在后续子串中寻找，否则j++
		if isPalindrome(s[i : j+1]) {
			dfs(s, j+1, append(cur, s[i:j+1]), result)
		}
	}
}

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	s := "abcd"
	result := partition(s)
	fmt.Println(result)
}
