package main

import (
	"fmt"
)

/*
要求：


解题思路：
到达(i,j)的最小值=min(到达(i-1,j)的最小值,到达(i,j-1)的最小值)+(i,j)

关键点：
数字在节点上,故相邻两节点间最小值即为两节点值之和(数字在线上时,两节点间最小值需考虑所有情况的最小值)

*/

func minPathSum(grid [][]int) int { // faster 99.64% less 100%.
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	res := grid
	for i := 1; i < len(grid); i++ {
		res[i][0] = grid[i][0] + res[i-1][0]
	}
	for i := 1; i < len(grid[0]); i++ {
		res[0][i] = grid[0][i] + res[0][i-1]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			res[i][j] = min(res[i-1][j], res[i][j-1]) + grid[i][j]
		}
	}
	return res[len(grid)-1][len(grid[0])-1]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	result := minPathSum(grid)
	fmt.Println(result)
}
