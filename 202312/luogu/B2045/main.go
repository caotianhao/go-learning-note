package main

import "fmt"

func main() {
	var a int
	_, _ = fmt.Scanf("%d", &a)
	if a == 1 || a == 3 || a == 5 {
		fmt.Printf("NO")
	} else {
		fmt.Printf("YES")
	}
}
