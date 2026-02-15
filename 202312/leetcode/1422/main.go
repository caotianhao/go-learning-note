package main

import (
	"fmt"
	"strings"
)

func maxScore(s string) int {
	maxN := -1
	for i := 1; i < len(s); i++ {
		t := strings.Count(s[:i], "0") + strings.Count(s[i:], "1")
		if t > maxN {
			maxN = t
		}
	}
	return maxN
}

func main() {
	fmt.Println(maxScore("011101"))
	fmt.Println(maxScore("11"))
}
