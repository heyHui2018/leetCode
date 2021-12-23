package main

import (
	"fmt"
	"math"
)

/*
要求：


解题思路：
持有，售出，休息
售出的收益等于持有加上当天价格
持有的收益等于max(前一天休息之后买入，前一天持有）
休息的收益等于max(前一天休息，前一天售出)


关键点：


*/

func maxProfit(prices []int) int { // faster 100% less 100%
	sold, rest, hold := 0, 0, math.MinInt32
	for _, v := range prices {
		preSold := sold
		sold = hold + v
		hold = max(hold, rest-v)
		rest = max(rest, preSold)
	}
	return max(rest, sold)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	result := maxProfit([]int{10, 9, 2, 5, 3, 7, 101, 18, 1})
	fmt.Println(result)
}
