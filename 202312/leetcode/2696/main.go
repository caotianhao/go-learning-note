package main

import "fmt"

func minLength(s string) int {
	stack := make([]byte, 0)
	for i := range s {
		if len(stack) != 0 &&
			(stack[len(stack)-1] == 'A' && s[i] == 'B' ||
				stack[len(stack)-1] == 'C' && s[i] == 'D') {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack)
}

func main() {
	fmt.Println(minLength("ABFCACDB"))
	fmt.Println(minLength("ACBBD"))
}
