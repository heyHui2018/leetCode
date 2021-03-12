package main

import (
	"fmt"
)

/*
要求：


解题思路：
因存在负数，故需记录最大及最小


关键点：


*/

func maxProduct(nums []int) int { // faster 76.92% less 100%
	result, minP, maxP := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		minP, maxP = min(minP*nums[i], maxP*nums[i], nums[i]), max(minP*nums[i], maxP*nums[i], nums[i])
		result = max(maxP, result)
	}
	return result
}

func max(a int, list ...int) int {
	temp := a
	for i := range list {
		if list[i] > temp {
			temp = list[i]
		}
	}
	return temp
}

func min(a int, list ...int) int {
	temp := a
	for i := range list {
		if list[i] < temp {
			temp = list[i]
		}
	}
	return temp
}

func main() {
	nums := []int{-1, -2}
	result := maxProduct(nums)
	fmt.Println(result)
}
