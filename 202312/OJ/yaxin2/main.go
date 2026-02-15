package main

import "fmt"

func isValid(s string) bool {
	stack := make([]byte, 0)
	for _, v := range s {
		if len(stack) != 0 && (stack[len(stack)-1] == '(' && v == ')' ||
			stack[len(stack)-1] == '[' && v == ']' ||
			stack[len(stack)-1] == '{' && v == '}') {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, byte(v))
		}
	}
	return len(stack) == 0
}

func main() {
	var s string
	_, _ = fmt.Scanf("%s", &s)
	fmt.Println(isValid(s))
}
