package main

import "fmt"

func main() {
	var a int
	_, _ = fmt.Scanf("%d", &a)
	if a >= 10 && a <= 99 {
		fmt.Printf("1")
	} else {
		fmt.Printf("0")
	}
}
