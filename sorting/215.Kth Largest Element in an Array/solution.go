package main

import (
	"math/rand"
	"time"
)

/*
要求：


解题思路：
归并


关键点：


*/

var list = []int{1, 3, 5, 7, 9, 0, 8, 6, 4, 2}
var s = 3

func main() {
	fastSelect(0, len(list)-1, list, len(list)-s)
}

func swap(a, b int, nums []int) {
	temp := nums[a]
	nums[a] = nums[b]
	nums[b] = temp
}

func fastSelect(left, right int, nums []int, k int) int { // faster 98.84% less 100%
	if left == right {
		return nums[left]
	}
	rand.Seed(time.Now().UnixNano())
	i := left + rand.Intn(right-left)
	pivot := nums[i]
	// 把pivot放到最右边
	swap(i, right, nums)
	// 把小于pivot的值放到左边
	curIndex := left
	for j := left; j < right; j++ {
		if nums[j] < pivot {
			swap(curIndex, j, nums)
			curIndex++
		}
	}
	swap(curIndex, right, nums) // 把pivot放回去
	// 3.判断curIndex和k的关系
	switch {
	case curIndex == k:
		return nums[curIndex]
	case curIndex < k:
		return fastSelect(curIndex+1, right, nums, k)
	default:
		return fastSelect(left, curIndex-1, nums, k)
	}
}
