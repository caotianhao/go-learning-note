package main

import (
	"fmt"
)

func removePalindromeSub(s string) int {
	l := len(s)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[l-1-i] {
			return 2
		}
	}
	return 1
}

func main() {
	fmt.Println(removePalindromeSub("abababababababababab"))
}
