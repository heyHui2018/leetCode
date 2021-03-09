package main

import (
	"fmt"
	"strings"
)

/*
要求：


解题思路：


关键点：


*/

func reverseWords(s string) string { // faster 100% less 13.52%
	list := strings.Split(s, " ")
	result := ""
	for i := len(list) - 1; i >= 0; i-- {
		if list[i] == "" {
			continue
		}
		result = result + list[i] + " "
	}
	return strings.TrimSpace(result)
}

func main() {
	s := "a good   example"
	result := reverseWords(s)
	fmt.Println(result)
}
