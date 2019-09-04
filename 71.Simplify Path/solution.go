package main

import (
	"fmt"
	"strings"
)

/*
要求：


解题思路：
//=/ ./=当前目录 ../=上层目录,用slice及top作为简易栈

关键点：


*/

func simplifyPath(path string) string { // faster 100% less 100%
	if len(path) == 0 {
		return ""
	}
	if path[0] == '/' {
		path = path[1:]
	}
	list := strings.Split(path, "/")
	temp := make([]string, len(list))
	top := 0
	for _, v := range list {
		switch v {
		case ".", "":
			break
		case "..":
			if top > 0 {
				top--
			}
		default:
			temp[top] = v
			top++
		}
	}
	res := "/"
	for k, v := range temp[:top] {
		res += v
		if k != len(temp[:top])-1 {
			res += "/"
		}
	}
	return res
}

func main() {
	path := "/a//b////c/d//././/.."
	result := simplifyPath(path)
	fmt.Println(result)
}
