package main

import (
	"fmt"
	"strconv"
)

/*
要求：


解题思路：
抽象成()?()，分治法，递归，得到左右两边的结果后，遍历计算


关键点：


*/

func diffWaysToCompute(expression string) []int { // faster 100% less 53.85%
	i, err := strconv.Atoi(expression)
	if err == nil {
		return []int{i}
	}
	var result []int
	for i := range expression {
		if expression[i] == '+' || expression[i] == '-' || expression[i] == '*' {
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])
			for j := range left {
				for k := range right {
					switch expression[i] {
					case '+':
						result = append(result, left[j]+right[k])
					case '-':
						result = append(result, left[j]-right[k])
					default:
						result = append(result, left[j]*right[k])
					}
				}
			}
		}
	}
	return result
}

func main() {
	result := diffWaysToCompute("2-1-1")
	fmt.Println(result)
}
