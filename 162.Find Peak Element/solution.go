package main

import (
	"fmt"
)

/*
要求：


解题思路：
nums[0]>nums[1]时，nums[0]即为peak，终点同理，故数组中必存在peak，此类寻找一般使用二分法


关键点：


*/

func findPeakElement(nums []int) int { // faster 100% less 100%
	if len(nums) == 0 {
		return 0
	}
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] < nums[mid+1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func main() {
	nums := []int{1, 2, 1, 3, 5, 6, 4}
	result := findPeakElement(nums)
	fmt.Println(result)
}
