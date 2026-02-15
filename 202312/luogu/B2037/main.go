package main

import "fmt"

func main() {
	var a int
	_, _ = fmt.Scanf("%d", &a)
	if a%2 == 0 {
		fmt.Printf("even")
	} else {
		fmt.Printf("odd")
	}
}
