package main

import "fmt"

func isCircularSentence(sentence string) bool {
	l := len(sentence)
	if sentence[0] != sentence[l-1] {
		return false
	}
	for i := 0; i < l; i++ {
		if sentence[i] == ' ' && sentence[i-1] != sentence[i+1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isCircularSentence("leetcode exercises sound delightful"))
	fmt.Println(isCircularSentence("aa"))
	fmt.Println(isCircularSentence("a b"))
}
