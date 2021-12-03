package main

import (
	"fmt"
	"sort"
)

/*
要求：


解题思路：
排序后分两组，从大到小从两组中先后取数


关键点：


*/

func wiggleSort(nums []int) { // faster 70.83% less 33.33%
	sort.Ints(nums)
	var min, max []int
	if len(nums)%2 == 0 {
		min = nums[:len(nums)/2]
		max = nums[len(nums)/2:]
	} else {
		min = nums[:len(nums)/2+1]
		max = nums[len(nums)/2+1:]
	}
	fmt.Println(min)
	fmt.Println(max)
	result := make([]int, len(nums))
	for i := 0; i < len(max); i++ {
		result[i*2] = min[len(min)-1-i]
		result[i*2+1] = max[len(max)-1-i]
	}
	if len(min) > len(max) {
		result[len(nums)-1] = min[0]
	}
	for i := range result {
		nums[i] = result[i]
	}
}

func main() {
	nums := []int{1, 5, 1, 1, 6, 4}
	wiggleSort(nums)
	fmt.Println(nums)
}
