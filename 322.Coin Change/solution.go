package main

import (
	"fmt"
)

/*
要求：


解题思路：
状态转移方程：dp[i] = min(dp[i-coins[j]]) + 1 (j取值范围为[0,len(coins)])


关键点：


*/

func coinChange(coins []int, amount int) int { // faster 29.17% less 5.56%
	if amount == 0 {
		return 0
	}

	var dp = make(map[int]int)
	// 初始化dp，组成每一种硬币所用的最少硬币数为1
	for i := 0; i < len(coins); i++ {
		dp[coins[i]] = 1
	}

	// 拟定一个最值，若最值无变化则说明不能凑成amount
	max := 9999

	for i := 1; i <= amount; i++ {
		min := max
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 && dp[i-coins[j]] < min { // 大于 最小的零钱 且 零钱数量更少，则更新零钱数量
				min = dp[i-coins[j]]
			}
		}
		dp[i] = min + 1
	}
	if dp[amount] > max {
		return -1
	}
	return dp[amount]
}

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	result := coinChange(coins, amount)
	fmt.Println(result)
}
