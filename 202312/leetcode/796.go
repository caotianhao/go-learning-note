package main

import "fmt"

func rotateString(s string, goal string) bool {
	l := len(s)
	for i := 0; i < l; i++ {
		s = rotate796(s)
		if s == goal {
			return true
		}
	}
	return false
}

func rotate796(s string) string {
	return s[1:] + string(s[0])
}

func main() {
	fmt.Println(rotateString("abcdef", "acb"))
}
