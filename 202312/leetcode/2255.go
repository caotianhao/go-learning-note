package main

import (
	"fmt"
	"strings"
)

func countPrefixes(words []string, s string) int {
	cnt, l := 0, len(words)
	for i := 0; i < l; i++ {
		if strings.HasPrefix(s, words[i]) {
			cnt++
		}
	}
	return cnt
}

func main() {
	words := []string{"a", "b", "c", "ab", "bc", "abc"}
	s := "abc"
	fmt.Println(countPrefixes(words, s))
}
