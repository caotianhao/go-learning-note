package main

import "fmt"

func minOperations1758(s string) int {
	l, zero, one, m, n := len(s), "", "", 0, 0
	for i := 0; i < l; i++ {
		if i%2 == 0 {
			zero += "0"
			one += "1"
		} else {
			zero += "1"
			one += "0"
		}
	}
	for i := 0; i < l; i++ {
		if s[i] != zero[i] {
			m++
		}
		if s[i] != one[i] {
			n++
		}
	}
	return min(m, n)
}

func main() {
	fmt.Println(minOperations1758("1010101010"))
	fmt.Println(minOperations1758("11"))
	fmt.Println(minOperations1758("00"))
}
