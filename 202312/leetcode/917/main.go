package main

import "fmt"

func reverseOnlyLetters(s string) string {
	l := len(s)
	allZiMu, other, rev := make([]byte, l), make([]byte, l), make([]byte, 0)
	for i := 0; i < l; i++ {
		if (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= 'a' && s[i] <= 'z') {
			allZiMu[i] = s[i]
		} else {
			other[i] = s[i]
		}
	}
	allZiMu = reverse917(allZiMu)
	for i := range allZiMu {
		if allZiMu[i] == 0 {
			continue
		}
		rev = append(rev, allZiMu[i])
	}
	for i := range other {
		if other[i] != 0 {
			rev = insert917(rev, i, other[i])
		}
	}
	return string(rev)
}

func reverse917(b []byte) []byte {
	l := len(b)
	for i := 0; i < l>>1; i++ {
		b[i], b[l-i-1] = b[l-i-1], b[i]
	}
	return b
}

func insert917(b []byte, index int, char byte) []byte {
	tmp := append([]byte{}, b[index:]...)
	b = append(append(b[:index], char), tmp...)
	return b
}

func main() {
	fmt.Println(reverseOnlyLetters("Test1ng-Leet=code-Q!")) //Qedo1ct-eeLg=ntse-T!
}
