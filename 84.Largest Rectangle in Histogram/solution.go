package main

import (
	"fmt"
)

/*
要求：


解题思路：
遍历数组,同时取left/right从i开始向两边遍历,遇到边界或值小于heights[i]时停止,并计算面积,迭代出最大的面积.需注意当left=0且heights[0]>heights[i]时,left不用+1

关键点：


*/

func largestRectangleArea(heights []int) int { // faster 6.35% less 100%
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
		fmt.Println("i = ", i, " left = ", left, " right = ", right, " area = ", area)
	}
	return temp
}

func main() {
	heights := []int{2, 1, 2}
	result := largestRectangleArea(heights)
	fmt.Println(result)
}
