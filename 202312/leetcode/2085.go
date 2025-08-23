package main

import "fmt"

func countWords(words1 []string, words2 []string) int {
	map1, map2, c := map[string]int{}, map[string]int{}, 0
	for _, v := range words1 {
		map1[v]++
	}
	for _, v := range words2 {
		map2[v]++
	}
	for _, v := range words2 {
		if cnt, ok := map1[v]; ok && cnt == 1 && map2[v] == 1 {
			c++
		}
	}
	return c
}

func main() {
	fmt.Println(countWords([]string{"leetcode", "is", "amazing", "as", "is"},
		[]string{"amazing", "leetcode", "is"}))
	fmt.Println(countWords([]string{"a", "ab"}, []string{"a", "a", "a", "ab"}))
}
