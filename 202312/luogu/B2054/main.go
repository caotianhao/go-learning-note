package main

import "fmt"

func main() {
	var n, age, sum int
	_, _ = fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d\n", &age)
		sum += age
	}
	fmt.Printf("%.2f", float64(sum)/float64(n))
}
