package main

import (
	"fmt"
	"strings"
)

func countVowelSubstrings(word string) (cnt int) {
	l := len(word)
	for i := 0; i < l; i++ {
		for j := i + 4; j < l; j++ {
			if help2062(word[i : j+1]) {
				cnt++
			}
		}
	}
	return
}

func help2062(str string) bool {
	m := map[string]struct{}{}
	for _, v := range str {
		m[string(v)] = struct{}{}
	}
	return strings.Contains(str, "a") && strings.Contains(str, "e") &&
		strings.Contains(str, "i") && strings.Contains(str, "o") &&
		strings.Contains(str, "u") && len(m) == 5
}

func main() {
	fmt.Println(countVowelSubstrings("cuaieuouac")) //7
}
