package main

import "fmt"

func truncateSentence(s string, k int) string {
	str, l, space := "", len(s), 0
	if k == 0 {
		return str
	}
	for i := 0; i < l; i++ {
		if s[i] == ' ' {
			space++
		}
	}
	if k > space {
		return s
	} else {
		space = 0
		for i := 0; i < l; i++ {
			if s[i] == ' ' {
				space++
			}
			if space == k {
				str += s[:i]
				break
			}
		}
	}
	return str
}

func main() {
	fmt.Println(truncateSentence("What is the solution to this problem", 4))
}
