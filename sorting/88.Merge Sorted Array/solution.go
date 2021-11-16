package main

import (
	"fmt"
)

/*
要求：


解题思路：
从后往前处理可避免频繁交换.遍历nums2,与nums1中的数从后往前比较,若nums2中的数较大,则将其放在nums1的未排序的最后并继续遍历,若nums1中的数较大,则将其放在nums1的未排序的最后,
并取nums1中的前一个数继续与nums2中的数比较直到nums2中的数被放到nums1的未排序的最后才开始下一轮遍历

关键点：


*/

func merge(nums1 []int, m int, nums2 []int, n int) { // faster 100% less 66.67%
	// p-nums1中待比较的位置,q-nums1中从后往前第一个未排序的位置
	p := m - 1
	q := len(nums1) - 1
	for i := n - 1; i >= 0; i-- {
		if p >= 0 {
			if nums2[i] >= nums1[p] {
				nums1[q] = nums2[i]
			} else {
				for p >= 0 && nums2[i] < nums1[p] {
					nums1[q] = nums1[p]
					nums1[p] = 0
					q--
					p--
				}
				nums1[q] = nums2[i]
			}
		} else {
			nums1[q] = nums2[i]
		}
		q--
	}
}

func main() {
	num1 := []int{2, 0}
	m := 1
	num2 := []int{1}
	n := 1
	merge(num1, m, num2, n)
	fmt.Println(num1)
}
