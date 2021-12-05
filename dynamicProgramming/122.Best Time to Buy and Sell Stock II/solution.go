package main

import (
	"fmt"
)

/*
要求：
as many transactions as you like

解题思路：
若当前价格高于前一天，则加入收益，即表示前一天持有

关键点：


*/

func maxProfit(prices []int) int { // faster 100% less 100%
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	result := maxProfit(prices)
	fmt.Println(result)
}
