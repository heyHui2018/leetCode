package main

import (
	"fmt"
)

/*
要求：


解题思路：
先取m(x1,y1)、n(x2,y2)两点，再遍历判断后续点是否在此线上
一般式：ax+by+c=0
两点式：(y-y1)/(y2-y1)=(x-x1)/(x2-x1) (x1≠x2，y1≠y2)
斜截式：y=kx+b
取两点式较简单

关键点：


*/

func maxPoints(points [][]int) int { // faster 98.4% less 100%
	if len(points) <= 2 {
		return len(points)
	}
	max := 0
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			count := 0
			for k := 0; k < len(points); k++ {
				if isSameLine(points[i], points[j], points[k]) {
					count++
				}
			}
			if count > max {
				max = count
			}

		}
	}
	return max
}

func isSameLine(a, b, c []int) bool {
	return (a[1]-b[1])*(c[0]-b[0]) == (c[1]-b[1])*(a[0]-b[0])
}

func main() {
	points := [][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}
	result := maxPoints(points)
	fmt.Println(result)
}
