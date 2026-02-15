package main

import "fmt"

func replaceDigits(s string) string {
	l, target := len(s), ""
	if l%2 == 0 {
		for i := 0; i < l; i += 2 {
			target += string(rune(s[i]))
			target += myShift(s[i], int(s[i+1]-48))
		}
	} else {
		for i := 0; i < l-1; i += 2 {
			target += string(rune(s[i]))
			target += myShift(s[i], int(s[i+1]-48))
		}
		target += string(s[l-1])
	}
	return target
}

func myShift(c byte, x int) string {
	return string(rune(int(c) + x))
}

func main() {
	s := "a1c1e1f"
	fmt.Println(replaceDigits(s))
}
