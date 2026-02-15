package main

import (
	"fmt"
)

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	myWord1, myWord2, sumWord1, sumWord2, flag := make([]string, 0), make([]string, 0), "", "", true
	for i := 0; i < len(word1); i++ {
		sumWord1 += word1[i]
	}
	for i := 0; i < len(word2); i++ {
		sumWord2 += word2[i]
	}
	myWord1 = append(myWord1, sumWord1)
	myWord2 = append(myWord2, sumWord2)
	if len(myWord1) != len(myWord2) {
		return false
	} else {
		for i := 0; i < len(myWord1); i++ {
			if myWord1[i] != myWord2[i] {
				flag = false
			}
		}
	}
	return flag
}

func main() {
	str1 := []string{"abc", "ddd"}
	str2 := []string{"abcqq", "dwedd"}
	fmt.Println(arrayStringsAreEqual(str1, str2))
}
