package main

import (
	"fmt"
)

/*
要求：


解题思路：
无序，缺少的数字 = 无丢失时的累加和 - 丢失时的累加和


关键点：


*/

func missingNumber(nums []int) int { // faster 98.59% less 100%
	total := 0
	for i := range nums {
		total += nums[i]
	}
	sum := len(nums) * (len(nums) + 1) / 2
	return sum - total
}

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
}
