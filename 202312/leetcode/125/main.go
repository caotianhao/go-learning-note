package main

import (
	"fmt"
	"strings"
)

func isPalindrome125(s string) bool {
	s = strings.ToLower(s)
	l, str := len(s), ""
	for i := 0; i < l; i++ {
		if s[i] >= 97 && s[i] <= 122 || s[i] >= 48 && s[i] <= 57 {
			str += string(s[i])
		}
	}
	ll := len(str)
	for i := 0; i < ll; i++ {
		if str[i] != str[ll-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome125("A man, a plan, a canal: Panama"))
}
