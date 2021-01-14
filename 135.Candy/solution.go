package main

import (
	"fmt"
)

/*
要求：


解题思路：
从两个方向分别计算，某个点的值为两个方向计算结果中的较大值

关键点：


*/

func candy(ratings []int) int { // faster 94.55% less 21.82%
	l := len(ratings)
	if l <= 1 {
		return l
	}

	left := make([]int, l)  // 由左往右计算
	right := make([]int, l) // 由右往左计算
	left[0] = 1
	right[l-1] = 1

	for i := 1; i < l; i++ {
		if ratings[i] > ratings[i-1] { // i比其左边的大
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}

		if ratings[l-i] < ratings[l-i-1] { // l-i-1比其右边的大
			right[l-i-1] = right[l-i] + 1
		} else {
			right[l-i-1] = 1
		}
	}
	result := 0
	for i := range ratings {
		result += max(left[i], right[i])
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	ratings := []int{1, 0, 2}
	result := candy(ratings)
	fmt.Println(result)
}
