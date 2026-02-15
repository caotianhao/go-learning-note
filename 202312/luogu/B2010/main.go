package main

import "fmt"

func main() {
	var a, b int
	_, _ = fmt.Scanf("%d %d", &a, &b)
	fmt.Printf("%d %d", a/b, a%b)
}
