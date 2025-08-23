package main

import "fmt"

func maxDepth(s string) int {
	l, stack, maxN := len(s), 0, -1
	if l == 0 || l == 1 {
		return 0
	}
	for _, v := range s {
		if v == '(' {
			stack++
		}
		if stack > maxN {
			maxN = stack
		}
		if v == ')' {
			stack--
		}
	}
	return maxN - stack
}

func main() {
	fmt.Println(maxDepth("(1+(2*3)+((8)/4))+1"))
	fmt.Println(maxDepth("(1)+((2))+(((3)))"))
	fmt.Println(maxDepth("()()"))
	fmt.Println(maxDepth("()(()())"))
}
