package main

import (
	"fmt"
	"strings"
)

/*
要求：


解题思路：


关键点：


*/

func isPalindrome(s string) bool { // faster 61.46% less 60.79%
	s = strings.ToLower(s)
	i := 0
	j := len(s) - 1
	for i < j {
		for i < j && !isChar(s[i]) {
			i++
		}
		for i < j && !isChar(s[j]) {
			j--
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isChar(c byte) bool {
	if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
		return true
	}
	return false
}

func main() {
	s := "aa"
	result := isPalindrome(s)
	fmt.Println(result)
}
