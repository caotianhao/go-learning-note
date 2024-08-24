package main

import "fmt"

func countAnagram0813(s string) (cnt int) {
	n := len(s)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if isAnagram0813(s[i : j+1]) {
				cnt++
			}
		}
	}
	return
}

func isAnagram0813(s string) bool {
	r := []rune(s)
	n := len(r)
	for i := 0; i < n/2; i++ {
		if r[i] != r[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var str string
	_, _ = fmt.Scanf("%s", &str)
	fmt.Println(countAnagram0813(str))
}
