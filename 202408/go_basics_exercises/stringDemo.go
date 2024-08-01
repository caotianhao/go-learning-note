package main

import (
	"fmt"
	"unicode/utf8"
)

func countRune(b []byte) int {
	return utf8.RuneCount(b)
}

func replace1(s string) string {
	return s[:3] + "abc" + s[4:]
}

func reverse1(s string) string {
	r := []rune(s)
	l := len(r)
	for i := 0; i < l/2; i++ {
		r[i], r[l-1-i] = r[l-1-i], r[i]
	}
	return string(r)
}

func main() {
	s1, s2 := "byDSWEW ggg ioxudwmn xy", "foobar"
	fmt.Println(countRune([]byte(s1)))
	fmt.Println(replace1(s1))
	fmt.Println(reverse1(s2))
}
