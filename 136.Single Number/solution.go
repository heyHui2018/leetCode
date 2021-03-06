package main

import (
	"fmt"
)

/*
要求：


解题思路：
a^a=0
b=a^a^b

关键点：


*/

func singleNumber(nums []int) int { // faster 91.3% less 40%
	result := 0
	for i := range nums {
		result ^= nums[i]
	}
	return result
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	result := singleNumber(nums)
	fmt.Println(result)
}
