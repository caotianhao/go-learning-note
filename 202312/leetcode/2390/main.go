package main

import "fmt"

func removeStars(s string) string {
	stack := make([]byte, 0)
	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			stack = append(stack, s[i])
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}

func main() {
	fmt.Println(removeStars("leet**cod*e"))
}
