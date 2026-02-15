package main

import "fmt"

func checkIfPangram(sentence string) bool {
	l := len(sentence)
	if l <= 25 {
		return false
	}
	charMap := make(map[string]int)
	for i := 0; i < l; i++ {
		charMap[string(sentence[i])]++
	}
	if len(charMap) == 26 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(checkIfPangram("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))
	fmt.Println(checkIfPangram("thequickbrownfoxjumpsoverthelazydog"))
}
