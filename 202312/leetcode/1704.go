package main

import "fmt"

func halvesAreAlike(s string) bool {
	l := len(s)
	s1, s2, cnt1, cnt2 := s[:l/2], s[l/2:], 0, 0
	for i := 0; i < l/2; i++ {
		if s1[i] == 'a' || s1[i] == 'e' || s1[i] == 'i' ||
			s1[i] == 'o' || s1[i] == 'u' || s1[i] == 'A' ||
			s1[i] == 'E' || s1[i] == 'I' || s1[i] == 'O' || s1[i] == 'U' {
			cnt1++
		}
	}
	for i := 0; i < l/2; i++ {
		if s2[i] == 'a' || s2[i] == 'e' || s2[i] == 'i' ||
			s2[i] == 'o' || s2[i] == 'u' || s2[i] == 'A' ||
			s2[i] == 'E' || s2[i] == 'I' || s2[i] == 'O' || s2[i] == 'U' {
			cnt2++
		}
	}
	return cnt1 == cnt2
}

func main() {
	fmt.Println(halvesAreAlike("book"))
}
