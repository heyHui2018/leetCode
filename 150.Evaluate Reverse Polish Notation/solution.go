package main

import (
	"fmt"
	"strconv"
)

/*
要求：


解题思路：
遍历，数字入栈，遇到符号则出栈两个数字进行计算再将结果入栈


关键点：


*/

func evalRPN(tokens []string) int { // faster 89.26% less 97.52%
	nums := make([]int, 0, len(tokens))
	for i := range tokens {
		if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
			b, a := nums[len(nums)-1], nums[len(nums)-2]
			nums = nums[:len(nums)-2]
			nums = append(nums, cal(a, b, tokens[i]))
		} else {
			temp, _ := strconv.Atoi(tokens[i])
			nums = append(nums, temp)
		}
	}
	return nums[0]
}

func cal(a, b int, opt string) int {
	switch opt {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	default:
		return a / b
	}
}

func main() {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	result := evalRPN(tokens)
	fmt.Println(result)
}
