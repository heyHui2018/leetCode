package main

import (
	"fmt"
)

/*
要求：


解题思路：
递归.遍历判断两个字符是否存在某一次分割,能使s1的某一子串与s2的某一子串互为搅拌字符,s1的另一子串与s2的另一子串互为搅拌字符,一旦存在则返回true,否则返回false
三维动态规划.

关键点：


*/

func isScramble(s1 string, s2 string) bool { // faster 100% less 100%
	// 相同即为互为搅拌字符
	if s1 == s2 {
		return true
	}
	// 校验包含字符是否一致
	rec := make([]int, 256)
	for i := 0; i < len(s1); i++ {
		rec[s1[i]]++
		rec[s2[i]]--
	}
	for i := range rec {
		if rec[i] != 0 {
			return false
		}
	}
	// 校验字符是否互为搅拌字符
	for i := 1; i < len(s1); i++ {
		// 当存在某一次分割使s1的左子串与s2的左子串互为搅拌字符,且s1的右子串与s2的右子串互为搅拌字符,true
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
			return true
		}
		// 当存在某一次分割使s1的左子串与s2的右子串互为搅拌字符,且s1的右子串与s2的左子串互为搅拌字符,true
		// s1[:i] i个字符, s2[len(s1)-i:] i个字符
		fmt.Println(s1[:i], s2[len(s1)-i:])
		if isScramble(s1[:i], s2[len(s1)-i:]) && isScramble(s1[i:], s2[:len(s1)-i]) {
			return true
		}
	}
	return false
}

func main() {
	s1 := "abcde"
	s2 := "caebd"
	result := isScramble(s1, s2)
	fmt.Println(result)
}
