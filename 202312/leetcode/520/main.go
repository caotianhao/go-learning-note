package main

import "fmt"

func detectCapitalUse(word string) bool {
	l := len(word)
	low, up := 0, 0
	for _, v := range word {
		if v >= 'a' && v <= 'z' {
			low++
		} else if v >= 'A' && v <= 'Z' {
			up++
		}
	}
	if low == l || up == l || low == l-1 && word[0] <= 'Z' && word[0] >= 'A' {
		return true
	}
	return false
}

func main() {
	fmt.Println(detectCapitalUse("USA"))
	fmt.Println(detectCapitalUse("U"))
	fmt.Println(detectCapitalUse("a"))
	fmt.Println(detectCapitalUse("aaa"))
	fmt.Println(detectCapitalUse("AaG"))
}
