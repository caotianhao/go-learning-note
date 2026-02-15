package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scanf("%d", &n)
	if n <= 5 {
		fmt.Printf("Local")
	} else {
		fmt.Printf("Luogu")
	}
}
