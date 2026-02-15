package main

import "fmt"

func main() {
	var a, b int
	_, _ = fmt.Scanf("%d\n%d", &a, &b)
	fmt.Printf("%.3f%%", float64(b)*100.0/float64(a))
}
