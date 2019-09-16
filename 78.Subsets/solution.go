package main

import (
	"fmt"
)

/*
要求：


解题思路：
在77题的基础上拓展,同样是回溯算法,77题仅需1-n自然数的长度=k的所有组合,而本题是传入数组nums的长度从0-len(nums)的各个长度的所有组合,修改方法c的start入参及nums替换原来的n即可

关键点：


*/

func subsets(nums []int) [][]int { // faster 100% less 50%
	res := [][]int{}
	for i := 0; i <= len(nums); i++ {
		c(&res, []int{}, 0, nums, i)
	}
	return res
}

func c(res *[][]int, coms []int, start int, nums []int, k int) {
	if k == 0 {
		tmp := []int{}
		tmp = append(tmp, coms...)
		*res = append(*res, tmp)
		return
	}
	for i := start; i < len(nums); i++ {
		coms = append(coms, nums[i])
		c(res, coms, i+1, nums, k-1)
		// 回溯
		coms = coms[:len(coms)-1]
	}
}

func main() {
	nums := []int{0}
	result := subsets(nums)
	for k := range result {
		fmt.Println(result[k])
	}
}
