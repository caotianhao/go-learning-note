package main

import "fmt"

func areOccurrencesEqual(s string) bool {
	l, strMap := len(s), make(map[string]int)
	for i := 0; i < l; i++ {
		strMap[string(s[i])]++
	}
	a := strMap[string(s[0])]
	for _, v := range strMap {
		if v != a {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(areOccurrencesEqual("abasbc"))
}
