package main

import "fmt"

func main() {
	var a, b int
	_, _ = fmt.Scanf("%d %d", &a, &b)
	if a >= 10 || b >= 20 {
		fmt.Printf("1")
	} else {
		fmt.Printf("0")
	}
}
