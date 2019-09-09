package main

import (
	"fmt"
)

/*
要求：


解题思路：
先根据每行的首位确定目标所要比较的行,随后二分法判断是否存在

关键点：


*/

func searchMatrix(matrix [][]int, target int) bool { // faster 96.98% less 100%
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	l := len(matrix[0])
	c := -1
	for k := range matrix {
		if matrix[k][l-1] >= target && target >= matrix[k][0] {
			c = k
			break
		}
	}
	if c == -1 {
		return false
	}
	left := 0
	right := l - 1
	for {
		if left == right || left == right-1 {
			if matrix[c][left] == target || matrix[c][right] == target {
				return true
			}
			break
		}
		switch {
		case matrix[c][(left+right)/2] > target:
			right -= (right - left) / 2
		case matrix[c][(left+right)/2] < target:
			left += (right - left) / 2
		default:
			return true
		}
	}
	return false
}

func main() {
	matrix := [][]int{{-9, -7, -7, -5, -3}, {-1, 0, 1, 3, 3}, {5, 7, 9, 11, 12}, {13, 14, 15, 16, 18}, {19, 21, 22, 22, 22}}
	target := -4
	result := searchMatrix(matrix, target)
	fmt.Println(result)
}
