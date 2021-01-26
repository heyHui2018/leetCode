package main

import (
	"fmt"
)

/*
要求：


解题思路：
没有memo记录时会因重复操作导致超时

关键点：


*/

func wordBreak(s string, wordDict []string) bool { // faster 100% less 8.56%
	wordMap := map[string]bool{}
	for _, v := range wordDict {
		wordMap[v] = true
	}
	memo := map[int]bool{}
	return canBreak(0, s, wordMap, memo)
}

func canBreak(start int, s string, wordMap map[string]bool, memo map[int]bool) bool {
	if start == len(s) {
		return true
	}
	if res, ok := memo[start]; ok {
		return res
	}
	for i := start + 1; i <= len(s); i++ {
		prefix := s[start:i]
		if wordMap[prefix] && canBreak(i, s, wordMap, memo) {
			memo[start] = true
			return true
		}
	}
	memo[start] = false
	return false
}

func main() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	result := wordBreak(s, wordDict)
	fmt.Println(result)
}
