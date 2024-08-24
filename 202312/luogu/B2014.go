package main

import "fmt"

func main() {
	var r float64
	const p = 3.14159
	_, _ = fmt.Scanf("%f", &r)
	fmt.Printf("%.4f %.4f %.4f", 2*r, 2*r*p, p*r*r)
}
