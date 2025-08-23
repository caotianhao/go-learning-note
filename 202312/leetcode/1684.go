package main

import "fmt"

func countConsistentStrings(allowed string, words []string) int {
	cnt, lw := 0, len(words)
	for i := 0; i < lw; i++ {
		if func1684(allowed, words[i]) {
			cnt++
		}
	}
	return cnt
}

func func1684(allowed, word string) bool {
	la, lw, cnt := len(allowed), len(word), 0
	for i := 0; i < lw; i++ {
		for j := 0; j < la; j++ {
			if word[i] == allowed[j] {
				cnt++
				break
			}
		}
	}
	if cnt == lw {
		return true
	} else {
		return false
	}
}

func main() {
	allowed := "ab"
	words := []string{"ad", "bd", "aaab", "baa", "badab"}
	fmt.Println(countConsistentStrings(allowed, words))
}
