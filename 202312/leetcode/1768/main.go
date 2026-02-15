package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	l1, l2, target := len(word1), len(word2), ""
	if l1 > l2 {
		for i := 0; i < l2; i++ {
			target += string(word1[i])
			target += string(word2[i])
		}
		for i := l2; i < l1; i++ {
			target += string(word1[i])
		}
		return target
	} else if l2 > l1 {
		for i := 0; i < l1; i++ {
			target += string(word1[i])
			target += string(word2[i])
		}
		for i := l1; i < l2; i++ {
			target += string(word2[i])
		}
		return target
	} else {
		for i := 0; i < l1; i++ {
			target += string(word1[i])
			target += string(word2[i])
		}
		return target
	}
}

func main() {
	fmt.Println(mergeAlternately("dictionary", "absolutely"))
}
