package main

import (
	"fmt"
)

/*
要求：


解题思路：
递归.第一行一个元素为1;第二行两个元素,第零个元素为1,第一个元素为其上一行的第一个元素与其左边元素之和即1...第n行第i(i>1)个元素等于第n-1行的第i-1个元素与第i个元素之和

关键点：


*/

func generate(numRows int) [][]int { // faster 100% less 50%
	res := make([][]int, numRows)
	for i := 1; i <= numRows; i++ {
		temp := make([]int, i)
		temp[0] = 1
		for j := 1; j < i-1; j++ {
			temp[j] = res[i-2][j-1] + res[i-2][j]
		}
		if i-1 > 0 {
			temp[i-1] = 1
		}
		res[i-1] = temp
	}
	return res
}

func main() {
	numRows := 0
	result := generate(numRows)
	for k := range result {
		fmt.Println(result[k])
	}
}
