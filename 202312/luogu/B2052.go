package main

import "fmt"

func main() {
	var a, b int
	var cal byte
	_, _ = fmt.Scanf("%d %d %c", &a, &b, &cal)
	if cal == '+' {
		fmt.Printf("%d", a+b)
	} else if cal == '-' {
		fmt.Printf("%d", a-b)
	} else if cal == '*' {
		fmt.Printf("%d", a*b)
	} else if cal == '/' {
		if b == 0 {
			fmt.Printf("Divided by zero!")
		} else {
			fmt.Printf("%d", a/b)
		}
	} else {
		fmt.Printf("Invalid operator!")
	}
}
