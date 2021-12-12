package main

import (
	"fmt"
)

/*
要求：


解题思路：
每10	个数，有 1个		个位 为1
每100	个数，有 10个		十位 为1
每1000	个数，有	100个	百位 为1

以算百位数字为例：
n=3141092	a=31410	b=92	百位上1的个数=31410*100
n=3141192	a=31411	b=92	百位上1的个数=31410*100+92+1
n=3141292	a=31412 b=92	百位上1的个数=31411*100

(a+8)/10*m + (a%10==1)*(b+1)

31410 = 314109 / 10，给31410多+8，从而使31412开始得到的记过会大1

关键点：


*/

func countDigitOne(n int) int { // faster 100% less 42.86%
	result := 0
	for m := 1; m <= n; m = m * 10 {
		a := n / m
		b := n % m
		result += (a + 8) / 10 * m
		if a%10 == 1 {
			result += b + 1
		}
	}
	return result
}

func main() {
	result := countDigitOne(1)
	fmt.Println(result)
}
