package main

import (
	"fmt"
)

/*
要求：


解题思路：
f(n)=f(n-1)+f(n-2),f(0)=1,f(1)=1

关键点：


*/

func climbStairs(n int) int { // faster 100% less 22.22%
	if n < 2 {
		return 1
	}
	rec := make([]int, n+1)
	rec[0], rec[1] = 1, 1
	// 用slice的遍历来代替函数的递归
	for i := 2; i <= n; i++ {
		rec[i] = rec[i-1] + rec[i-2]
	}
	return rec[n]
}

func main() {
	n := 45
	result := climbStairs(n)
	fmt.Println(result)
}
