package main

import "fmt"

func main() {
	var a, b, n int
	_, _ = fmt.Scanf("%d %d %d", &a, &b, &n)
	diff := b - a
	fmt.Printf("%d", a+diff*(n-1))
}
