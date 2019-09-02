package main

import (
	"fmt"
)

/*
要求：


解题思路：
牛顿法

关键点：


*/

func mySqrt(x int) int { // faster 100% less 83.33%.
	res := x
	// 牛顿法求平方根
	for res*res > x {
		res = (res + x/res) / 2
	}
	return res
}

func main() {
	x := 188
	result := mySqrt(x)
	fmt.Println(result)
}
