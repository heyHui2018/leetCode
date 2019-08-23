package main

import (
	"fmt"
)

/*
要求：


解题思路：
到达(i,j)的路径数=到达(i-1,j)的路径数+到达(i,j-1)的路径数

关键点：


*/

func uniquePaths(m int, n int) int { // faster 100% less 100%.
	// 二维数组初始化时长度不能为变量,而题中note指出m/n最大100,故在此设置为100
	path := [100][100]int{}
	// 到达第0列和第0行的路径数=1
	for i := 0; i < m; i++ {
		path[i][0] = 1
	}
	for i := 0; i < n; i++ {
		path[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			path[i][j] = path[i-1][j] + path[i][j-1]
		}
	}
	return path[m-1][n-1]
}

func main() {
	result := uniquePaths(6, 7)
	fmt.Println(result)
}
