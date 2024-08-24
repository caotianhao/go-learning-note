package main

import "fmt"

func main() {
	var n int
	var num, sum float64
	_, _ = fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%f", &num)
		sum += num
	}
	fmt.Printf("%.6f", sum/float64(n))
}
