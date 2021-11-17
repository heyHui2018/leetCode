package main

import (
	"fmt"
)

/*
要求：


解题思路：
用map存储遍历过的数字
排序后遍历看是否有相邻大小一致的


关键点：


*/

func containsDuplicate(nums []int) bool { // faster 73.8% less 33.2%
	m := make(map[int]bool)
	for i := range nums {
		if m[nums[i]] {
			return true
		}
		m[nums[i]] = true
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 7}
	fmt.Println(containsDuplicate(nums))
}
