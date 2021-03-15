package main

import (
	"fmt"
)

/*
要求：


解题思路：
双指针


关键点：


*/

func findMin(nums []int) int { // faster 91.78% less 21.92%
	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}

	i, j := 0, len(nums)-1
	for i < j {
		if nums[i] > nums[i+1] {
			return nums[i+1]
		}

		if nums[j-1] > nums[j] {
			return nums[j]
		}

		i++
		j--
	}

	return nums[i]
}

func main() {
	nums := []int{2, 2, 2, 0, 1}
	result := findMin(nums)
	fmt.Println(result)
}
