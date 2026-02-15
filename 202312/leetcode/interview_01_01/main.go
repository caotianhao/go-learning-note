package main

import "fmt"

func isUnique(astr string) bool {
	var char [26]int
	for _, v := range astr {
		char[v-'a']++
	}
	for _, v := range char {
		if v > 1 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isUnique(""))
}
