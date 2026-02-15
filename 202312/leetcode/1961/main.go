package main

import "fmt"

func isPrefixString(s string, words []string) bool {
	ss := ""
	for _, v := range words {
		if s == ss {
			return true
		}
		ss += v
	}
	return s == ss
}

func main() {
	fmt.Println(isPrefixString("iloveleetcode", []string{"i", "love", "leetcode", "apples"}))
	fmt.Println(isPrefixString("iloveleetcode", []string{"apples", "i", "love", "leetcode"}))
	fmt.Println(isPrefixString("z", []string{"z"}))
}
