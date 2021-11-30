package main

import (
	"fmt"
)

/*
要求：
sort in-place
can not use the library's sort function

解题思路：
取i,k标识首尾,取j从头扫描,遇到0时i位置与j位置的数交换且i、j均++,
遇到1时j继续扫描,遇到2时j位置与k位置的数交换且k--,
直到j>k时返回,此时0-i位置的数均为0,i-j位置额数均为1,j往后的均为2

关键点：


*/

func sortColors(nums []int) { // faster 100% less 100%
	i, j, k := 0, 0, len(nums)-1
	// for 循环中， nums[i:j] 中始终全是 1
	// 循环结束后，
	// nums[:i] 中全是 0
	// nums[j:] 中全是 2
	for j <= k {
		switch nums[j] {
		case 0:
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		case 1:
			j++
		case 2:
			nums[j], nums[k] = nums[k], nums[j]
			k--
		}
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}
