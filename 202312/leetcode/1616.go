package main

import "fmt"

func checkPalindromeFormation(a string, b string) bool {
	if isPalindrome1616(a) || isPalindrome1616(b) {
		return true
	}
	l := len(a)
	for i := 0; i < l; i++ {
		if isPalindrome1616(a[:i] + b[i:]) {
			return true
		}
	}
	for i := 0; i < l; i++ {
		if isPalindrome1616(b[:i] + a[i:]) {
			return true
		}
	}
	return false
}

func isPalindrome1616(s string) bool {
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-1-i] {
			return false
		}
	}
	return true
}

func main() {
	//fmt.Println(checkPalindromeFormation("a", "b"))           //t
	//fmt.Println(checkPalindromeFormation("abdef", "fecab"))   //t
	//fmt.Println(checkPalindromeFormation("ulacfd", "jizalu")) //t
	fmt.Println(checkPalindromeFormation("abda", "acmc")) //f
	//fmt.Println(isPalindrome1616(""))
}
