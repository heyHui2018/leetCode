package main

import (
	"fmt"
)

/*
要求：


解题思路：
丑数仅有因子2、3、5，则丑数*2/3/5=丑数，


关键点：


*/

func nthUglyNumber(n int) int { // faster 100% less 29.67%
	res := []int{1}

	// 分别记录第一个乘以2、3、5后大于当前
	index2 := 0
	index3 := 0
	index5 := 0

	for len(res) < n {
		m := min(res[index2]*2, min(res[index3]*3, res[index5]*5))

		if m == res[index2]*2 {
			index2++
		}
		if m == res[index3]*3 {
			index3++
		}
		if m == res[index5]*5 {
			index5++
		}
		res = append(res, m)
	}

	return res[len(res)-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	result := nthUglyNumber(10)
	fmt.Println(result)
}
