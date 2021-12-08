package main

import (
	"fmt"
)

/*
要求：


解题思路：
深度优先遍历，发现1后递归遍历，遍历完一块联通的1后将其置0，再往下遍历


关键点：


*/

func numIslands(grid [][]byte) int { // faster 24.16% less 59.92%
	if len(grid) == 0 {
		return 0
	}
	result := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				result++
			}
		}
	}
	return result
}

func dfs(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) {
		return
	}
	if grid[i][j] == '1' {
		grid[i][j] = '0'
		dfs(grid, i+1, j)
		dfs(grid, i-1, j)
		dfs(grid, i, j+1)
		dfs(grid, i, j-1)
	}
}

func main() {
	result := numIslands([][]byte{})
	fmt.Println(result)
}
