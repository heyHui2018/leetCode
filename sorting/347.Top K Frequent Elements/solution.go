package main

import (
	"fmt"
	"sort"
)

/*
要求：


解题思路：
map统计各数字出现频次，二维数组统计各频次有哪些数字，从大到小输出即可


关键点：


*/

func topKFrequent(nums []int, k int) []int { // faster 87.06% less 90.36%
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]]++
	}
	var f []int
	for k := range m {
		f = append(f, k)
	}
	sort.Slice(f, func(a, b int) bool {
		return m[f[a]] > m[f[b]]
	})
	return f[:k]
}

func main() {
	nums := []int{4, 1, -1, 2, -1, 2, 3}
	fmt.Println(topKFrequent(nums, 2))
}
