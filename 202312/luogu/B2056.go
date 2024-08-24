package main

import "fmt"

func main() {
	var n, num, sum int
	_, _ = fmt.Scan(&n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&num)
		sum += num
	}
	fmt.Printf("%d %.5f\n", sum, float64(sum)/float64(n))
}
