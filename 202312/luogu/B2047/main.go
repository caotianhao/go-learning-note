package main

import "fmt"

func main() {
	var a float32
	_, _ = fmt.Scanf("%f", &a)
	if 0 <= a && a < 5 {
		fmt.Printf("%.3f", 2.5-a)
	} else if 5 <= a && a < 10 {
		fmt.Printf("%.3f", 2-1.5*(a-3)*(a-3))
	} else {
		fmt.Printf("%.3f", (a-3)/2)
	}
}
