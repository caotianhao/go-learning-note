package main

import (
	"fmt"
)

func minStringValue(s string, k int) int {
	stack := make([]int32, 0)
	for _, v := range s {
		if len(stack) == 0 {
			stack = append(stack, v)
		} else {
			if stack[len(stack)-1] == v {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, v)
			}
		}
	}
	return len(stack) - k<<1
}

func main() {
	var (
		n int
		k int
		s string
	)
	_, _ = fmt.Scanf("%d %d", &n, &k)
	_, _ = fmt.Scanf("%s", &s)
	fmt.Println(minStringValue(s, k))
}
