package main

import (
	"fmt"
)

/*
要求：


解题思路：
从后向前遍历,相加,注意进位,string到int再到string的转换通过string(int(a[ka]-'0')+up+'0')实现

关键点：


*/

func addBinary(a string, b string) string { // faster 100% less 50%.
	ka, kb := len(a)-1, len(b)-1
	up := 0
	res := ""
	for {
		switch {
		case ka < 0 && kb < 0:
			if up > 0 {
				res = "1" + res
			}
			return res
		case ka >= 0 && kb < 0:
			if int(a[ka]-'0')+up > 1 {
				res = "0" + res
				up = 1
			} else {
				res = string(int(a[ka]-'0')+up+'0') + res
				up = 0
			}
			ka--
		case ka < 0 && kb >= 0:
			if int(b[kb]-'0')+up > 1 {
				res = "0" + res
				up = 1
			} else {
				res = string(int(b[kb]-'0')+up+'0') + res
				up = 0
			}
			kb--
		case ka >= 0 && kb >= 0:
			if int(a[ka]-'0')+int(b[kb]-'0')+up > 1 {
				res = string(int(a[ka]-'0')+int(b[kb]-'0')+up-2+'0') + res
				up = 1
			} else {
				res = string(int(a[ka]-'0')+int(b[kb]-'0')+up+'0') + res
				up = 0
			}
			ka--
			kb--
		}
	}
	return res
}

func main() {
	a := "11010"
	b := "1011"
	result := addBinary(a, b)
	fmt.Println(result)
}
