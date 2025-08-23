package main

import "fmt"

func checkAlmostEquivalent(word1 string, word2 string) bool {
	var slice1, slice2 [26]int
	for _, v := range word1 {
		slice1[v-97]++
	}
	for _, v := range word2 {
		slice2[v-97]++
	}
	for i := 0; i < 26; i++ {
		if slice1[i]-slice2[i] > 3 || slice2[i]-slice1[i] > 3 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(checkAlmostEquivalent("aaa", "bcb"))
}
