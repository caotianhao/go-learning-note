package main

import (
	"fmt"
	"strings"
)

func reorderSpaces(text string) (res string) {
	space := 0
	for _, v := range text {
		if v == ' ' {
			space++
		}
	}
	fi := strings.Fields(text)
	word := len(fi)
	sp := ""
	for i := 0; i < space; i++ {
		sp += " "
	}
	if word == 1 {
		return fi[0] + sp
	}
	whole, remain := space/(word-1), space%(word-1)
	mid, re := "", ""
	for i := 0; i < whole; i++ {
		mid += " "
	}
	for i := 0; i < remain; i++ {
		re += " "
	}
	for i := 0; i < word-1; i++ {
		res += fi[i] + mid
	}
	res += fi[word-1] + re
	return
}

func main() {
	fmt.Println(reorderSpaces("  this   is  a sentence "))
	fmt.Println(reorderSpaces("aaa"))
	fmt.Println(reorderSpaces("a b   "))
	fmt.Println(reorderSpaces("         ehh"))
}
