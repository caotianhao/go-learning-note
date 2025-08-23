package main

import "fmt"

func removeDuplicates1047(s string) string {
	stack1047 := make([]byte, 0)
	for i := range s {
		if len(stack1047) > 0 && stack1047[len(stack1047)-1] == s[i] {
			stack1047 = stack1047[:len(stack1047)-1]
		} else {
			stack1047 = append(stack1047, s[i])
		}
	}
	return string(stack1047)
}

func main() {
	fmt.Println(removeDuplicates1047("abbaca"))
}
