package main

import (
	"fmt"
)

/*
要求：


解题思路：
dp[i][j]表示以(i,j)为右下角的正方形的最大边长
0 1 1 1 0		0 1 1 1 0
1 1 1 1 0		1 1 2 2 0
0 1 1 1 1	=>	0 1 2 3 1
0 1 1 1 1		0 1 2 3 2
0 0 1 1 1		0 0 1 2 3
dp[i][j] = min(dp[i-1][j],dp[i][j-1],dp[i-1],[j-1])+1


关键点：


*/

func maximalSquare(matrix [][]byte) int { // faster 50.93% less 22.98%
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
	}
	max := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '0' {
				continue
			}
			dp[i][j] = int(matrix[i][j] - '0')
			if i > 0 && j > 0 {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			}
			if dp[i][j] > max {
				max = dp[i][j]
			}
		}
	}
	return max * max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// matrix := [][]byte{
	// 	{'1', '0', '1', '0', '0'},
	// 	{'1', '0', '1', '1', '1'},
	// 	{'1', '1', '1', '1', '1'},
	// 	{'1', '0', '0', '1', '0'},
	// }
	matrix := [][]byte{
		{'0', '1'},
		{'1', '0'},
	}
	result := maximalSquare(matrix)
	fmt.Println(result)
}
