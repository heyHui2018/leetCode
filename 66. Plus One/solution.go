package main

import (
	"fmt"
)

/*
要求：


解题思路：
从后往前遍历,无需进位时可直接返回,出现类似99需在最高位进位时,需重新申请数组并将第一位置1

关键点：


*/

func plusOne(digits []int) []int { // faster 100% less 50%.
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			// 第一位=9,即最高位需进一
			if i == 0 {
				new := make([]int, len(digits)+1)
				new[0] = 1
				return new
			}
			continue
		}
		digits[i] += 1
		return digits
	}
	return digits
}

func main() {
	digits := []int{9, 9, 9}
	result := plusOne(digits)
	fmt.Println(result)
}
