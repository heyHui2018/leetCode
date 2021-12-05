package main

import (
	"fmt"
)

/*
要求：


解题思路：
统计矩阵每一列在每一行上,从这一列开始往后连续的1的个数(即遇到0停止计数)从而形成数组,随后采用84题中计算数组能形成的最大矩形的算法来计算此数组,最终迭代出矩阵能形成的最大矩形

关键点：


*/

func maximalRectangle(matrix [][]byte) int { // faster 76.14% less 20%
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	tempArea := 0
	area := 0
	for j := 0; j < len(matrix[0]); j++ {
		var list []int
		for top := 0; top < len(matrix); top++ {
			count := 0
			for left := j; left < len(matrix[0]); left++ {
				if matrix[top][left] == '0' {
					break
				}
				count++
			}
			list = append(list, count)
		}
		area = maxArea(list)
		if tempArea < area {
			tempArea = area
		}
	}

	return tempArea
}

func maxArea(heights []int) int {
	temp := 0
	for i := 0; i < len(heights); i++ {
		left, right := i, i
		for {
			if left > 0 && heights[left] >= heights[i] {
				left--
			} else if right < len(heights) && heights[right] >= heights[i] {
				right++
			} else {
				if left < i {
					if left != 0 || heights[left] < heights[i] {
						left = left + 1
					}
				}
				if right > i {
					right = right - 1
				}
				break
			}
		}
		area := heights[i] * (right - left + 1)
		if area > temp {
			temp = area
		}
	}
	return temp
}

func main() {
	matrix := [][]byte{[]byte("10100"), []byte("10111"), []byte("11111"), []byte("10010")}
	result := maximalRectangle(matrix)
	fmt.Println(result)
}
