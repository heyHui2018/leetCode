package main

import (
	"fmt"
)

/*
要求：


解题思路：
深度优先遍历，当前点的增长路径最大值=比当前点值大的相邻点的增长路径最大值+1

关键点：


*/

var memo [][]int

// dfs的四个方向
var dirs = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func longestIncreasingPath(matrix [][]int) int { // faster 25.71% less 49.52%
	if len(matrix) == 0 {
		return 0
	}
	res := 0
	memo = make([][]int, len(matrix))
	visited := make([][]bool, len(matrix))
	for i := range memo {
		memo[i] = make([]int, len(matrix[0]))
		visited[i] = make([]bool, len(matrix[0]))
	}
	for i := range matrix {
		for j := range matrix[i] {
			res = max(res, dfs(matrix, i, j, visited))
		}
	}
	return res
}

func dfs(matrix [][]int, r, c int, visited [][]bool) int {
	if memo[r][c] > 0 { // 若当前点能在记录中查到，则直接返回查到的结果
		return memo[r][c]
	}

	res := 0

	visited[r][c] = true // 将当前点标记为已访问

	for _, dir := range dirs { // 循环4个方向
		// 下个点坐标
		row := dir[0] + r
		col := dir[1] + c

		// 下个坐标在矩阵内且当前路径未被访问且当前点小于下个点
		if row >= 0 && col >= 0 && row < len(matrix) && col < len(matrix[0]) && !visited[row][col] && matrix[r][c] < matrix[row][col] {
			res = max(res, dfs(matrix, row, col, visited))
		}
	}
	visited[r][c] = false // 恢复当前点未被访问
	memo[r][c] = res + 1
	return res + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	matrix := [][]int{{1}, {2}, {5}}
	result := longestIncreasingPath(matrix)
	fmt.Println(result)
}
