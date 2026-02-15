package main

import "fmt"

func isValid(s string) bool {
	stack20 := make([]byte, 0)
	for i := range s {
		if len(stack20) > 0 &&
			(stack20[len(stack20)-1] == s[i]-1 || stack20[len(stack20)-1] == s[i]-2) {
			stack20 = stack20[:len(stack20)-1]
		} else {
			stack20 = append(stack20, s[i])
		}
	}
	return len(stack20) == 0
}

func main() {
	fmt.Println(isValid("()[]{(}"))
}
