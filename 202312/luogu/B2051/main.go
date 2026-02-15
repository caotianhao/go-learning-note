package main

import "fmt"

func main() {
	var a, b int
	_, _ = fmt.Scanf("%d %d", &a, &b)
	if a >= -1 && a <= 1 && b <= 1 && b >= -1 {
		fmt.Printf("yes")
	} else {
		fmt.Printf("no")
	}
}
