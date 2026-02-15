package main

import "fmt"

func main() {
	var f, c float64
	_, _ = fmt.Scanf("%f", &f)
	c = 5.0 * (f - 32) / 9.0
	fmt.Printf("%.5f", c)
}
