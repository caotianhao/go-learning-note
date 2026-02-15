package main

import (
	"fmt"
	"strings"
)

func reverseStr(s string, k int) string {
	l := len(s)
	before, cnt, remain := make([]string, 0), l/(2*k), l%(2*k)
	ans := make([]string, 0)
	for i := 0; i < cnt; i++ {
		before = append(before, s[2*i*k:(i+1)*2*k])
	}
	for i := 0; i < len(before); i++ {
		tmp, str := []rune(before[i]), ""
		for j := 0; j < k/2; j++ {
			tmp[j], tmp[k-1-j] = tmp[k-1-j], tmp[j]
		}
		for i := 0; i < len(tmp); i++ {
			str += string(tmp[i])
		}
		ans = append(ans, str)
	}
	if remain >= k {
		tmp, str := []rune(s[l-remain:l-remain+k]), ""
		for j := 0; j < k/2; j++ {
			tmp[j], tmp[k-1-j] = tmp[k-1-j], tmp[j]
		}
		for i := 0; i < len(tmp); i++ {
			str += string(tmp[i])
		}
		ans = append(ans, str)
		ans = append(ans, s[l-remain+k:])
	} else {
		tmp, str := []rune(s[l-remain:]), ""
		for j := 0; j < len(tmp)/2; j++ {
			tmp[j], tmp[len(tmp)-1-j] = tmp[len(tmp)-1-j], tmp[j]
		}
		for i := 0; i < len(tmp); i++ {
			str += string(tmp[i])
		}
		ans = append(ans, str)
	}
	return strings.Join(ans, "")
}

func main() {
	//fmt.Println(reverseStr("abcd", 2))
	fmt.Println(reverseStr("abcdefg", 2))
	fmt.Println(reverseStr("abcdefg", 8))
	fmt.Println(reverseStr("abcdefghi", 3))
}
