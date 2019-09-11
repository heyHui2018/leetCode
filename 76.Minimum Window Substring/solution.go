package main

import (
	"fmt"
)

/*
要求：
in complexity O(n)

解题思路：
统计t中出现的字符及数量,遍历s,当发现需要的字符时,定为窗口的左侧,继续遍历直到窗口包括了t中的所有字符,
此时开始收缩左侧窗口直到刚好包括t中所有字符,记为一次满足记录，与已有记录比较,更小时覆盖记录,最终返回最小的记录

关键点：


*/

func minWindow(s string, t string) string { // faster 100% less 100%
	have := [128]int{}
	need := [128]int{}
	// 存放ASCII码0-128的数组,need[t[i]]++可继续字符的出现次数
	for i := range t {
		need[t[i]]++
	}
	size := len(s)
	total := len(t)
	min := size + 1
	res := ""

	for i, j, count := 0, 0, 0; j < size; j++ {
		if have[s[j]] < need[s[j]] {
			// 出现需要的字符
			count++
		}
		have[s[j]]++

		for i <= j && have[s[i]] > need[s[i]] {
			// 缩小窗口
			have[s[i]]--
			i++
		}

		width := j - i + 1
		if count == total && min > width {
			min = width
			res = s[i : j+1]
		}
	}
	return res
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	result := minWindow(s, t)
	fmt.Println(result)
}
