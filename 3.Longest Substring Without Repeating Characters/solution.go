package main

import (
	"fmt"
)

/*
要求:

尝试：

学习：
用数组存储字符出现的位置，当出现重复时，计算上次到本次出现的长度并与记录比较/替换，随后从上一次出现的位置往后计算

关键点：

*/

func lengthOfLongestSubstring(s string) int {
	// location[s[i]] == j 表示：s中第i个字符串,上次出现在s的j位置,所以在s[j+1:i]中没有s[i]
	// location[s[i]] == -1 表示：s[i]未在s中出现
	location := [256]int{} // 只有256长是因为假定输入的字符串只有ASCII字符
	for i := range location {
		location[i] = -1 // 先设置所有的字符都未出现
	}

	maxLen, left := 0, 0

	for i := 0; i < len(s); i++ {
		if location[s[i]] >= left {
			// 说明s[i]已经在s[left:i+1]中重复
			left = location[s[i]] + 1 // 从上一次出现的位置往后计算
		} else if i+1-left > maxLen {
			// 不重复的情况下，每一次遍历都更新maxLen
			maxLen = i + 1 - left
		}
		location[s[i]] = i
	}
	return maxLen
}

func main() {
	s := "abcb"
	result := lengthOfLongestSubstring(s)
	fmt.Println(result)
}
