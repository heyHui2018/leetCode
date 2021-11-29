package main

import (
	"fmt"
)

/*
要求：


解题思路：
遍历s记数各字母，遍历t在s的记录上减

关键点：


*/

func isAnagram(s string, t string) bool { // faster 28.88% less 19.52%
	if len(s) != len(t) {
		return false
	}
	m := make(map[string]int)
	for _, v := range s {
		m[string(v)]++
	}
	for _, v := range t {
		m[string(v)]--
		if m[string(v)] < 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("1", "2"))
}
