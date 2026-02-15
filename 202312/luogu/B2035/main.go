package main

import "fmt"

func main() {
	var a int
	_, _ = fmt.Scanf("%d", &a)
	if a == 0 {
		fmt.Printf("zero")
	} else if a < 0 {
		fmt.Printf("negative")
	} else {
		fmt.Printf("positive")
	}
}
