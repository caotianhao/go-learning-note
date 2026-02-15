package main

import "fmt"

func main() {
	var num float32
	_, _ = fmt.Scanf("%f\n", &num)
	fmt.Printf("%.3f", num)
}
