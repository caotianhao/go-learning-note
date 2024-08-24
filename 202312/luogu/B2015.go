package main

import "fmt"

func main() {
	var r1, r2 float64
	_, _ = fmt.Scanf("%f %f", &r1, &r2)
	fmt.Printf("%.2f", 1.0/((1.0/r1)+(1.0/r2)))
}
