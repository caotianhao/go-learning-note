package main

import "fmt"

func minimizedStringLength(s string) int {
	help := map[rune]struct{}{}
	for _, v := range s {
		help[v] = struct{}{}
	}
	return len(help)
}

func main() {
	fmt.Println(minimizedStringLength("abccc"))
}
