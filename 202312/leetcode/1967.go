package main

import "fmt"

func numOfStrings(patterns []string, word string) int {
	cnt, lp := 0, len(patterns)
	for i := 0; i < lp; i++ {
		if isSubstr(word, patterns[i]) {
			cnt++
		}
	}
	return cnt
}

func isSubstr(all string, sub string) bool {
	la, ls, flag := len(all), len(sub), false
	for i := 0; i < ls; i++ {
		for j := 0; j < la; j++ {
			if sub[i] == all[j] {
				if j+ls == la && sub == all[j:la] {
					flag = true
				}
				if j+ls < la && sub == all[j:j+ls] {
					flag = true
				}
			}
		}
	}
	return flag
}

func main() {
	patterns := []string{"a", "abc", "bc", "d"}
	word := "abc"
	fmt.Println(numOfStrings(patterns, word))
}
