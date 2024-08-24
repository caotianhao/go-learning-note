package main

import "fmt"

func main() {
	var a int
	_, _ = fmt.Scanf("%d", &a)
	if !(a%3 == 0 && a%5 == 0) {
		fmt.Printf("NO")
	} else {
		fmt.Printf("YES")
	}
}
