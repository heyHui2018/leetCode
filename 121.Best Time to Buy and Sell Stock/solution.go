package main

import (
	"fmt"
)

/*
要求：
at most one transaction

解题思路：
从第二位往后遍历，取当前位与前一位的差，并与之前的差累加，当差小于0时置0，表示从当天开始重新持有，当差大于max时赋值给max并继续遍历。

关键点：


*/

func maxProfit(prices []int) int { // faster 100% less 100%
	max := 0
	temp := 0
	for i := 1; i < len(prices); i++ {
		temp += prices[i] - prices[i-1]
		if temp < 0 {
			temp = 0
		}
		if max < temp {
			max = temp
		}
	}
	return max
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	result := maxProfit(prices)
	fmt.Println(result)
}
