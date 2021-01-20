package main

import (
	"fmt"
)

/*
要求：


解题思路：
三次，需要两个标志位
第 1 次出现时，ones = 5, twos = 0,  ones 记录这个数字
第 2 次出现时，ones = 0, twos = 5， twos 记录了这个数字
第 3 次出现时，ones = 0, twos = 0， 都清空了，可以去处理其他数字了
所以，如果某个数字出现了 1 次，ones 会记录，出现了 2 次，twos 会记录
公式：ones 要等于零， 而这时 twos 是 True ， 所以再 & 一个 twos 的非就可以
	ones = (ones xor i) & ~twos
	twos = (twos xor i) & ~ones

关键点：


*/

func singleNumber(nums []int) int { // faster 98.63% less 100%
	ones, twos := 0, 0
	for i := range nums {
		ones = (ones ^ nums[i]) & ^twos
		twos = (twos ^ nums[i]) & ^ones
	}
	return ones
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	result := singleNumber(nums)
	fmt.Println(result)
}
