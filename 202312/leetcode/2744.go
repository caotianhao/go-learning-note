package main

import "fmt"

func maximumNumberOfStringPairs(words []string) int {
	cnt := 0
	l := len(words)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if words[i] == reverse2744(words[j]) {
				cnt++
			}
		}
	}
	return cnt
}

func reverse2744(s string) string {
	r := []rune(s)
	l := len(r)
	for i := 0; i < l/2; i++ {
		r[i], r[l-i-1] = r[l-i-1], r[i]
	}
	return string(r)
}

func main() {
	fmt.Println(maximumNumberOfStringPairs([]string{"cd", "ac", "dc", "ca", "zz"}))
}
