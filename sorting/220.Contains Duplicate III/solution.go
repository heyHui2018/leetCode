package main

import (
	"fmt"
)

/*
要求：


解题思路：
索引差值k以内存在某个数差值小于t，返回false
双指针
无需判断index=i的两边，因为"比较i和i+k" = "比较i+k和i"


关键点：


*/

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool { // faster 34.25% less 60.27%
	if len(nums) == 0 {
		return false
	}
	for i := 1; i < len(nums); i++ {
		j := 0
		if i-k > 0 {
			j = i - k
		}
		for ; j < i; j++ {
			if abs(nums[i]-nums[j]) <= t {
				return true
			}
		}
	}
	return false
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func main() {
	nums := []int{1, 0, 1, 1}
	fmt.Println(containsNearbyAlmostDuplicate(nums, 1, 2))
}
