package main

import (
	"fmt"
)

/*
要求：


解题思路：
遍历nums1存入map，遍历nums2与map比较得出交集


关键点：


*/

func intersection(nums1 []int, nums2 []int) []int { // faster 100% less 51.69%
	m := make(map[int]struct{})
	for _, v := range nums1 {
		m[v] = struct{}{}
	}
	var result []int
	mm := make(map[int]struct{})
	for _, v := range nums2 {
		_, ok1 := m[v]
		_, ok2 := mm[v]
		if ok1 && !ok2 {
			mm[v] = struct{}{}
			result = append(result, v)
			continue
		}
	}
	return result
}

func main() {
	nums := []int{4, 1, -1, 2, -1, 2, 3}
	fmt.Println(intersection(nums, nums))
}
