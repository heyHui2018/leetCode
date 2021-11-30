package main

import (
	"fmt"
)

/*
要求：
in-place

解题思路：
33题的延伸,可直接使用二分法,33题中的数组无重复,故可直接比较中间值与最右值,
当中间值较小时,右半段有序,中间值较大时,左半段有序(原数组升序排列),
本题数组有重复,会出现中间值等于最右值的情况,可将最右值左移直到与中间值不相等为止

关键点：


*/

func search(nums []int, target int) bool { // faster 90.76% less 100%
	if len(nums) == 0 {
		return false
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}
		if nums[mid] < nums[right] { // 中间值小于最右值,右边有序
			if nums[mid] < target && nums[right] >= target { // 中间值小于目标且最右值大于目标,取右半部分继续二分即left = mid + 1
				left = mid + 1
			} else { // 取左半部分继续二分即right = mid - 1
				right = mid - 1
			}
			continue
		}
		if nums[mid] > nums[right] { // 中间值大于最右值,左边有序
			if nums[left] <= target && nums[mid] > target { // 最左值小于目标且中间值大于目标,取左半部分继续二分即right = mid - 1
				right = mid - 1
			} else { // 取右半部分继续二分即left = mid + 1
				left = mid + 1
			}
			continue
		}
		// 中间值等于最右值,right--
		right--
	}
	return false
}

func main() {
	nums := []int{2, 5, 6, 0, 0, 1, 2}
	target := 3
	result := search(nums, target)
	fmt.Println(result)
}
