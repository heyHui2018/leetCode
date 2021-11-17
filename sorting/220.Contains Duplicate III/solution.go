package main

import (
	"fmt"
)

/*
要求：


解题思路：


关键点：


*/

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool { // faster 98.84% less 100%
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
	fmt.Println(containsNearbyAlmostDuplicate(nums, 1, 1))
}
