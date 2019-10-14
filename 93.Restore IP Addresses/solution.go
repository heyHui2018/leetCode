package main

import (
	"fmt"
	"strconv"
)

/*
要求：


解题思路：
递归.从ip的第0段开始处理,for循环处理每一段ip的长度可能性,循环中判断长度是否满足要求,值是否合法等,若成立则递归下一段ip段,否则continue至下一个for循环

关键点：
3位时需<=255且最高位不能为0,2位时最高位不能为0

*/

func restoreIpAddresses1(s string) []string { // faster 5.95% less 100%
	var res []string
	if len(s) < 4 || len(s) > 12 {
		return res
	}
	r(s, 0, "", &res)
	return res
}

func r(s string, n int, str string, res *[]string) {
	if n == 3 && isValid(s) {
		*res = append(*res, str+s)
	}
	for i := 1; i < 4 && i < len(s); i++ {
		if isValid(string(s[:i])) {
			r(string(s[i:]), n+1, str+s[:i]+".", res)
		}
	}
}

func isValid(s string) bool {
	if s == "" || len(s) > 3 || (len(s) > 1 && s[0] == '0') {
		return false
	}
	num, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return num <= 255 && num >= 0
}

func restoreIpAddresses(s string) []string { // faster 100% less 100%
	res := []string{}
	n := len(s)
	if n < 4 || n > 12 {
		return res
	}
	combination := make([]string, 4)
	var dfs func(int, int)
	// index ip地址的段数
	dfs = func(idx, begin int) {
		if idx == 3 {
			temp := s[begin:]
			if isValid(temp) {
				combination[3] = temp
				res = append(res, fmt.Sprintf("%s.%s.%s.%s", combination[0], combination[1], combination[2], combination[3]))
			}
			return
		}
		// 剩余ip段可能的最大长度
		maxRemain := 3 * (3 - idx)
		for end := begin + 1; end <= n-(3-idx); end++ {
			// 当前情况下的长度总和过短
			if end+maxRemain < n {
				continue
			}
			// 当前ip段长度过长
			if end-begin > 3 {
				break
			}
			temp := s[begin:end]
			if isValid(temp) {
				combination[idx] = temp
				dfs(idx+1, end)
			}
		}
	}
	dfs(0, 0)
	return res
}

func main() {
	s := "25525511135"
	result := restoreIpAddresses(s)
	fmt.Println(result)
}
