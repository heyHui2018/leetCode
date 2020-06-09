package main

import (
	"fmt"
)

/*
要求:

尝试：

学习：
合并后取中位数(长度为奇时,中间值;为偶时,中间两值平均数)

关键点：

*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 { // faster 96.34% less 40%
	lenMis, i := len(nums1), 0
	lenNjs, j := len(nums2), 0
	res := make([]int, lenMis+lenNjs)

	for k := 0; k < lenMis+lenNjs; k++ {
		if i == lenMis ||
			(i < lenMis && j < lenNjs && nums1[i] > nums2[j]) {
			res[k] = nums2[j]
			j++
			continue
		}

		if j == lenNjs ||
			(i < lenMis && j < lenNjs && nums1[i] <= nums2[j]) {
			res[k] = nums1[i]
			i++
		}
	}
	l := len(res)
	if l%2 == 0 {
		return float64(res[l/2]+res[l/2-1]) / 2.0
	}

	return float64(res[l/2])
}

func main() {
	var a = []int{}
	var b = []int{1}
	result := findMedianSortedArrays(a, b)
	fmt.Println(result)
}
