package main

import "fmt"

func main() {
	var num float64
	_, _ = fmt.Scanf("%f", &num)
	fmt.Printf("%f\n%.5f\n%e\n%.6g", num, num, num, num)
}
