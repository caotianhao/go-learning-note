package main

import "fmt"

func main() {
	var a byte
	_, _ = fmt.Scanf("%c", &a)
	if a%2 == 0 {
		fmt.Printf("NO")
	} else {
		fmt.Printf("YES")
	}
}
