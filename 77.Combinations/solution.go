package main

import (
	"fmt"
)

/*
要求：


解题思路：
回溯算法,当某一次形成长度=k的组合后,回溯到之前k-1的长度并继续迭代形成新的组合,当此次k-1的迭代完成后,回溯到长度=k-2继续迭代

关键点：


*/

func combine(n int, k int) [][]int { // faster 65.64% less 100%
	var res [][]int
	c(&res, []int{}, 1, n, k)
	return res
}

func c(res *[][]int, coms []int, start int, n int, k int) {
	if k == 0 {
		tmp := []int{}
		tmp = append(tmp, coms...)
		*res = append(*res, tmp)
		return
	}
	for i := start; i <= n; i++ {
		coms = append(coms, i)
		fmt.Println("1111111111 coms = ", coms)
		c(res, coms, i+1, n, k-1)
		// 回溯
		coms = coms[:len(coms)-1]
		fmt.Println("22222 coms = ", coms)
	}
}

func main() {
	n := 4
	k := 2
	result := combine(n, k)
	fmt.Println(result)
}
