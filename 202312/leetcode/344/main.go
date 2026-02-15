package main

import "fmt"

//func reverseString(s []byte) {
//	l := len(s)
//	for i := 0; i < l/2; i++ {
//		s[i], s[l-i-1] = s[l-i-1], s[i]
//	}
//}

func reverseString(s []byte) {
	l := len(s)
	for i := 0; i < l/2; i++ {
		s[i], s[l-1-i] = s[l-i-1], s[i]
	}
}

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	fmt.Println(s)
}
