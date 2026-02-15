package main

import "fmt"

func main() {
	var a, b, c int
	_, _ = fmt.Scanf("%d %d %d", &a, &b, &c)
	fmt.Printf("%8d %8d %8d", a, b, c)
}
