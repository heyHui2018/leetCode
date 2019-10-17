package main

import (
	"fmt"
)

/*
要求：


解题思路：
n=0,res=1;n=1,res=1;n=2,res=2;n=3;res=5;res(n)=res(0)*res(n-1)+res(1)*res(n-2)+...+res(n-1)*res(0)卡特兰数

关键点：


*/

func numTrees(n int) int { // faster 100% less 100%
	list := make([]int, n+1)
	list[0] = 1
	list[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			list[i] += list[j] * list[i-j-1]
		}
	}
	fmt.Println(list)
	return list[n]
}

func main() {
	n := 3
	result := numTrees(n)
	fmt.Println(result)
}
