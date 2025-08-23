package main

import "fmt"

// 如果一个字符串不含有任何重复字符，我们称这个字符串为 好 字符串。
func countGoodSubstrings(s string) int {
	l, cnt := len(s), 0
	if l < 3 {
		return 0
	}
	for i := 0; i <= l-3; i++ {
		if isGood1876(s[i : i+3]) {
			cnt++
		}
	}
	return cnt
}

func isGood1876(s string) bool {
	map1876 := map[byte]int{}
	map1876[s[0]] = 0
	map1876[s[1]] = 1
	map1876[s[2]] = 2
	if len(map1876) == 3 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(countGoodSubstrings("ai"))
}
