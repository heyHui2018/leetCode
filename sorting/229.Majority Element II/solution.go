package main

import (
	"fmt"
	"math"
)

/*
要求：


解题思路：
摩尔投票算法
大于1/3的数只可能有2个

关键点：


*/

func majorityElement(nums []int) []int { // faster 27.78% less 18.06%
	if len(nums) == 0 {
		return []int{}
	}
	if len(nums) == 1 {
		return nums
	}
	n1, n2 := nums[0], math.MaxInt32
	c1, c2 := 1, 0
	for i := 1; i < len(nums); i++ {
		switch {
		case nums[i] == n1:
			c1++
		case nums[i] == n2:
			c2++
		case c1 == 0:
			n1 = nums[i]
			c1 = 1
		case c2 == 0:
			n2 = nums[i]
			c2 = 1
		default:
			c1--
			c2--
		}
	}
	t1, t2 := 0, 0
	for i := 0; i < len(nums); i++ {
		switch nums[i] {
		case n1:
			t1++
		case n2:
			t2++
		}
	}
	var result []int
	if t1 > len(nums)/3 {
		result = append(result, n1)
	}
	if t2 > len(nums)/3 {
		result = append(result, n2)
	}
	return result
}

func main() {
	nums := []int{2, 2, 1, 3}
	fmt.Println(majorityElement(nums))
}
