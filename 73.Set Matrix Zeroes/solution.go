package main

import (
	"fmt"
)

/*
要求：
Do it in-place

解题思路：
遍历矩阵,当发现某行出现0时,在遍历完该行后将此行全置为0,同时将出现0的列存入map,最后遍历map将对应的列置为0

关键点：


*/

func setZeroes(matrix [][]int) { // faster 85.78% less 100%
	columnMap := map[int]int{}
	for i := range matrix {
		flag := false
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				flag = true
				columnMap[j] = 1
			}
		}
		if flag {
			for k := range matrix[i] {
				matrix[i][k] = 0
			}
		}
	}
	for m := range columnMap {
		for i := range matrix {
			matrix[i][m] = 0
		}
	}
}

func main() {
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix)
	for k := range matrix {
		fmt.Println(matrix[k])
	}
}
