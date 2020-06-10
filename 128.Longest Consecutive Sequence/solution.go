package main

import (
	"fmt"
)

/*
要求：


解题思路：
先将array转为map，随后遍历map，判断是否存在当前数字的前一位，若存在则继续遍历，否则开始往后寻找所有连续的序列。
此法仅需处理各个序列的起点。

关键点：


*/

func longestConsecutive(nums []int) int { // faster 98。35% less 63.7%
	if len(nums) <= 1 {
		return len(nums)
	}
	max := 0
	intMap := make(map[int]bool)
	for _, v := range nums {
		intMap[v] = true
	}
	for _, v := range nums {
		if !intMap[v-1] {
			temp := 0
			for intMap[v] {
				temp++
				v++
			}
			if temp > max {
				max = temp
			}
		}
	}
	return max
}

func main() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	result := longestConsecutive(nums)
	fmt.Println(result)
}
