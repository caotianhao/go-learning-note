package main

import "fmt"

func main() {
	var a, b int
	_, _ = fmt.Scanf("%d %d", &a, &b)
	fmt.Printf("%.9f", float64(a)/float64(b))
}
