package main

import "fmt"

func maxLengthBetweenEqualCharacters(s string) (ans int) {
	l, flag := len(s), true
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if s[i] == s[j] {
				ans = max(ans, j-i-1)
				flag = false
			}
		}
	}
	if flag {
		return -1
	}
	return ans
}

func main() {
	fmt.Println(maxLengthBetweenEqualCharacters("abca"))
	fmt.Println(maxLengthBetweenEqualCharacters("document"))
	fmt.Println(maxLengthBetweenEqualCharacters("a"))
}
