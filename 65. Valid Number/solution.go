package main

import (
	"fmt"
)

/*
要求：


解题思路：
去除前后空格,随后遍历校验字符,小数点个数,是否以e或.e开头

关键点：


*/

func isNumber(s string) bool { // faster 100% less 100%.
	// 去除前后空格
	s = trim(s)
	// 校验是否实数
	return isRealNum(s)
}

func trim(s string) string {
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	j := len(s) - 1
	for i < j && s[j] == ' ' {
		j--
	}
	return s[i : j+1]
}

func isRealNum(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] == '-' || s[0] == '+' {
		return isNonnegReal(s[1:], false)
	}
	return isNonnegReal(s, false)
}

func isNonnegReal(s string, hasDot bool) bool {
	if len(s) == 0 {
		return false
	}
	for i, c := range s {
		switch {
		case '0' <= c && c <= '9':
			continue
		case c == '.':
			// 已存在小数点
			if hasDot {
				return false
			}
			// 以小数点结尾
			if i == len(s)-1 && i != 0 {
				return true
			}
			// 2.e3 true 但 .e3 false
			if i < len(s)-1 && s[i+1] == 'e' {
				return i != 0 && isInteger(s[i+2:])
			}
			return isNonnegReal(s[i+1:], true)
		case c == 'e':
			// e开头
			if i == 0 {
				return false
			}
			return isInteger(s[i+1:])
		default:
			return false
		}
	}
	return true
}

func isInteger(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] == '-' || s[0] == '+' {
		return isNonnegativeInteger(s[1:])
	}
	return isNonnegativeInteger(s)
}

func isNonnegativeInteger(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if c < '0' || '9' < c {
			return false
		}
	}
	return true
}

func main() {
	s := "0"
	result := isNumber(s)
	fmt.Println(result)
}
