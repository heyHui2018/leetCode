package main

import (
	"fmt"
	"strconv"
)

/*
要求：


解题思路：
若能，gas之和 一定 >= cost之和

关键点：


*/

func canCompleteCircuit(gas []int, cost []int) int { // faster 97.47% less 31.65%
	remains, debts, start := 0, 0, 0
	for i, g := range gas { // 从头遍历
		remains += g - cost[i] // 判断当前节点本身的盈亏
		if remains < 0 {       // >0表示有走到下一节点的能力，累计remains，<0时将start指向下一节点并重置remains
			start = i + 1    // i+1处重新开始
			debts += remains // 总
			remains = 0      // 置0
		}
		fmt.Println(i, "+"+strconv.Itoa(g), "-"+strconv.Itoa(cost[i]), "剩余"+strconv.Itoa(remains), "总"+strconv.Itoa(debts))
	}
	if debts+remains < 0 {
		// 最终油量，如无法全部偿还前期欠缺的油量，则无法跑完一圈
		return -1
	}
	return start
}

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	result := canCompleteCircuit(gas, cost)
	fmt.Println(result)
}
