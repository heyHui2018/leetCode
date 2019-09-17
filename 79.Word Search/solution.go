package main

import (
	"fmt"
)

/*
要求：


解题思路：
遍历board寻找与word中第一个字符匹配的字符board[i][j],找到后在该字符周围寻找与word中第二个字符匹配的字符,以此类推直至全部找到,否则返回false

关键点：


*/

func exist(board [][]byte, word string) bool { // faster 100% less 100%
	if len(word) == 0 {
		return true
	}
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if word[0] != board[i][j] {
				continue
			}
			if e(board, word, i, j) {
				return true
			}
		}
	}
	return false
}

func e(board [][]byte, word string, i, j int) bool {
	fmt.Println("111111111111 i = ", i, " j = ", j, " board[i][j] = ", string(board[i][j]))
	if len(word) == 0 { // 已匹配完成,true
		return true
	}
	if board[i][j] != word[0] { // 当前字符与word中第一个字符不匹配,false
		return false
	}
	if len(word) == 1 { // 当前字符与word中第一个字符匹配且word仅有一个字符,true
		return true
	}
	board[i][j] = '#' // 修改当前字符防止被多次使用
	if i > 0 && e(board, word[1:], i-1, j) { // 匹配上一行的字符
		return true
	}
	if j > 0 && e(board, word[1:], i, j-1) { // 匹配左侧的字符
		return true
	}
	if i < len(board)-1 && e(board, word[1:], i+1, j) { // 匹配下一行的字符
		return true
	}
	if j < len(board[0])-1 && e(board, word[1:], i, j+1) { // 匹配右侧的字符
		return true
	}
	board[i][j] = word[0] // 将字符复原
	return false
}

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "SEE"
	result := exist(board, word)
	fmt.Println(result)
}
