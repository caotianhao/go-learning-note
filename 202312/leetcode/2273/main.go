package main

import (
	"fmt"
	"reflect"
)

func removeAnagrams(words []string) []string {
	flag := true
	for i := 1; i < len(words); i++ {
		if isAnagrams2273(words[i-1], words[i]) {
			flag = false
		}
	}
	if flag {
		return words
	}
	for i := 1; i < len(words); i++ {
		if isAnagrams2273(words[i-1], words[i]) {
			words = delete2273(words, i)
		}
	}
	return removeAnagrams(words)
}

func isAnagrams2273(str1, str2 string) bool {
	m1, m2 := map[byte]int{}, map[byte]int{}
	for _, v := range str1 {
		m1[byte(v)]++
	}
	for _, v := range str2 {
		m2[byte(v)]++
	}
	return reflect.DeepEqual(m1, m2)
}

func delete2273(strs []string, ind int) []string {
	slice2273 := make([]string, 0)
	for i, v := range strs {
		if i == ind {
			continue
		}
		slice2273 = append(slice2273, v)
	}
	return slice2273
}

func main() {
	fmt.Println(removeAnagrams([]string{"abba", "baba", "bbaa", "cd", "cd"}))
}
