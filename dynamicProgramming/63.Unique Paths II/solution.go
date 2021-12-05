package main

import (
	"fmt"
)

/*
要求：


解题思路：
到达(i,j)的路径数=到达(i-1,j)的路径数+到达(i,j-1)的路径数,当(i,j)为障碍时,此点路径数=0

关键点：
初始化第一行及第一列的路径数,发现障碍时,后续位置路径数均为0

*/

func uniquePathsWithObstacles(obstacleGrid [][]int) int { // faster 100% less 100%.
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}
	res := [100][100]int{}
	for k := range obstacleGrid {
		if obstacleGrid[k][0] != 0 {
			break
		}
		res[k][0] = 1
	}
	for k := range obstacleGrid[0] {
		if obstacleGrid[0][k] != 0 {
			break
		}
		res[0][k] = 1
	}
	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[i]); j++ {
			if obstacleGrid[i][j] != 0 {
				continue
			}
			res[i][j] = res[i-1][j] + res[i][j-1]
		}
	}
	return res[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}

func main() {
	obstacleGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	result := uniquePathsWithObstacles(obstacleGrid)
	fmt.Println(result)
}
