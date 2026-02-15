package main

import (
	"fmt"
	"strings"
)

func prefixCount(words []string, pref string) int {
	cnt, l := 0, len(words)
	for i := 0; i < l; i++ {
		if strings.HasPrefix(words[i], pref) {
			cnt++
		}
	}
	return cnt
}

func main() {
	words := []string{"pay", "attention", "practice", "attend"}
	pref := "at"
	fmt.Println(prefixCount(words, pref))
}
