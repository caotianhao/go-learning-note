package main

import "fmt"

func firstPalindrome(words []string) string {
	l := len(words)
	for i := 0; i < l; i++ {
		if isPalindrome2(words[i]) {
			return words[i]
		}
	}
	return ""
}

func isPalindrome2(s string) bool {
	l, flag := len(s), true
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-1-i] {
			flag = false
		}
	}
	return flag
}

func main() {
	fmt.Println(firstPalindrome([]string{"abc", "car", "ada", "racecar", "cool"}))
}
