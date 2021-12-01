package main

import (
	"fmt"
	"sort"
)

/*
要求：


解题思路：
排序，从前往后找 值 大于等于 剩余元素数 的序号，结果为剩余元素数


关键点：


*/

func hIndex(citations []int) int { // faster 100% less 100%
	sort.Ints(citations)
	for i := range citations {
		if citations[i] >= len(citations)-i {
			return len(citations) - i
		}
	}
	return 0
}

func main() {
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5}))
}
