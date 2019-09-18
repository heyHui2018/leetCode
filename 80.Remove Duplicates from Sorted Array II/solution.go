package main

import (
	"fmt"
)

/*
要求：
in-place

解题思路：
取i表示处理完成的数的下标,取j=i往后遍历,当nums[i-2]!=nums[j]即发现一个新的数时,将i位置的数与j位置的数交换并将i+1,此时j的值也+1,继续遍历

关键点：


*/

func removeDuplicates(nums []int) int { // faster 93.30% less 100%
	i := 2
	for j := i; j < len(nums); j++ {
		if nums[i-2] != nums[j] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

func main() {
	nums := []int{1, 2, 2, 2, 2, 3}
	result := removeDuplicates(nums)
	fmt.Println(result)
}
