package main

import "fmt"

func isAcronym(words []string, s string) bool {
	ss := ""
	for _, w := range words {
		ss += string(w[0])
	}
	return ss == s
}

func main() {
	fmt.Println(isAcronym([]string{"never", "gonna", "give", "up", "on", "you"}, "ngguoy"))
}
