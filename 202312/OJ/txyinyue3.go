package main

import "fmt"

const mod1 = int(1e9 + 7)

func cntOfMethod(t string) int {
	if t[:2] == "aa" {
		return 2 + mod1
	}
	if t[0] == 'd' && t[1] == 'u' {
		return 4
	}
	return 3
}

func main() {
	fmt.Println(cntOfMethod("ababa"))
}
