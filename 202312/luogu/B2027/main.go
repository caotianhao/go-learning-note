package main

import "fmt"

func main() {
	var a int
	_, _ = fmt.Scanf("%d", &a)
	fmt.Printf("%.5f", 4.0*3.14*float64(a)*float64(a)*float64(a)/3.0)
}
