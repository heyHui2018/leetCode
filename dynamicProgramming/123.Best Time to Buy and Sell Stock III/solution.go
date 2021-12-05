package main

import (
	"fmt"
)

/*
要求：
at most two transactions

解题思路：
最多交易两次，即寻找i，使[0,i]和[i,len(prices)-1]两个区间中的能取得的最大值之和最大。
设f(i)为[0,i]的最大值数组，从前往后遍历；g(i)为[i,len(prices)-1]的最大值数组，从后往前遍历。

关键点：


*/

func maxProfit(prices []int) int { // faster 93.97% less 100%
	if len(prices) < 2 {
		return 0
	}
	f := make([]int, len(prices))
	g := make([]int, len(prices))

	least := prices[0]
	for i := 1; i < len(prices); i++ {
		least = min(least, prices[i])
		f[i] = max(f[i-1], prices[i]-least)
	}

	most := prices[len(prices)-1]
	for i := len(prices) - 2; i >= 0; i-- {
		most = max(most, prices[i])
		g[i] = max(g[i], most-prices[i])
	}

	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		maxProfit = max(maxProfit, f[i]+g[i])
	}
	return maxProfit
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	result := maxProfit(prices)
	fmt.Println(result)
}
