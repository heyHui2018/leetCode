package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*
要求：


解题思路：
在78题的基础上拓展,同样是回溯算法,78题仅需传入数组nums的长度从0-len(nums)的各个长度的所有组合,这里的nums是无重复的,修改方法s多传入一个map用于过滤即可

关键点：


*/

func subsetsWithDup(nums []int) [][]int { // faster 80.99% less 100%
	res := [][]int{}
	m := make(map[string]int)
	for i := 0; i <= len(nums); i++ {
		s(i, &res, []int{}, nums, 0, m)
	}
	return res
}

// k-子集长度,start-nums中剩余待处理的开始索引
func s(k int, res *[][]int, coms []int, nums []int, start int, m map[string]int) {
	if k == 0 {
		tmp := []int{}
		tmp = append(tmp, coms...)
		if len(tmp) > 0 {
			sort.Ints(tmp)
			num := ""
			for j := 0; j < len(tmp); j++ {
				num += strconv.Itoa(tmp[j])
			}
			if _, ok := m[num]; ok {
				return
			}
			m[num] = 1
		}
		*res = append(*res, tmp)
		return
	}
	for i := start; i < len(nums); i++ {
		coms = append(coms, nums[i])
		s(k-1, res, coms, nums, i+1, m)
		// 回溯
		coms = coms[:len(coms)-1]
	}
}

func main() {
	nums := []int{4, 1, 0}
	result := subsetsWithDup(nums)
	fmt.Println(result)
}
