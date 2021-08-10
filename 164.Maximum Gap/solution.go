package main

import (
	"fmt"
)

/*
要求：


解题思路：
排序并保存最大差值
线性的时间和空间复杂度，可使用桶排序
当前排序后相邻元素的最大与最小值取值于"当前桶内的最小值与前一个非空桶的最大值"

关键点：


*/

func maximumGap(nums []int) int { // faster 29.17% less 5.56%
	if len(nums) < 2 {
		return 0
	}

	min := nums[0]
	max := nums[0]

	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		} else if nums[i] < min {
			min = nums[i]
		}
	}

	diff := max - min

	section := diff/len(nums) + 1

	mMin, mMax := map[int]int{}, map[int]int{}

	for _, n := range nums {
		loc := (n - min) / section

		// 记录&更新每个桶的最大最小值
		if v, ok := mMin[loc]; !ok || n < v {
			mMin[loc] = n
		}

		if v, ok := mMax[loc]; !ok || n > v {
			mMax[loc] = n
		}
	}

	maxGap := 0
	lastNotEmptyLoc := 0
	// 索引每个桶中元素，当前桶最小值与上一个非空桶的最大值
	for i := 1; i < len(nums); i++ {
		if min, ok := mMin[i]; ok {
			t := min - mMax[lastNotEmptyLoc]
			if t > maxGap {
				maxGap = t
			}
			lastNotEmptyLoc = i
		}
	}
	return maxGap
}

func main() {
	nums := []int{3, 6, 9, 1}
	result := maximumGap(nums)
	fmt.Println(result)
}
