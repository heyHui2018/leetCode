package main

import (
	"fmt"
)

/*
要求：


解题思路：


关键点：


*/

func findMin(nums []int) int { // faster 100% less 18.69%
	i := 1
	for i < len(nums) && nums[i-1] < nums[i] {
		i++
	}
	// nums[0]可能就是min，此时i将=len(nums)，故需i=i%len(nums)
	return nums[i%len(nums)]
}

func main() {
	nums := []int{11, 13, 15, 17}
	result := findMin(nums)
	fmt.Println(result)
}
