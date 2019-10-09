package main

import (
	"fmt"
)

/*
要求：


解题思路：
直接排列:以二进制的0值为第零项,第一项改变最右边的位元,第二项改变右起第一个为1的位元的左边位元,第三、四项方法同第一、二项,以此类推.
镜射排列:n位元的格雷码可以从n-1位元的格雷码以上下镜射后加上新位元的方式得到,利用以下规则构造:
1、1位格雷码有两个码字;2、(n+1)位格雷码中的前2^n个码字等于n位格雷码的码字,按顺序书写,加前缀0;3、(n+1)位格雷码中的后2^n个码字等于n位格雷码的码字,按逆序书写,加前缀1

关键点：
格雷码:任意两个相邻的代码只有一位二进制数不同,最大数与最小数之间也仅一位数不同.

*/

func grayCode(n int) []int { // faster 710% less 100%
	return recur(n, 1, []int{0})
}

func recur(n, base int, nums []int) []int {
	if n == 0 {
		return nums
	}

	length := len(nums)
	temp := make([]int, length)
	for i := range nums {
		temp[length-i-1] = nums[i] + base
	}

	return recur(n-1, base*2, append(nums, temp...))
}

func main() {
	n := 3
	result := grayCode(n)
	fmt.Println(result)
}
