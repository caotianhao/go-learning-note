package main

import (
	"fmt"
	"sort"
)

func abs2095(a float64) float64 {
	if a > 0 {
		return a
	}
	return -a
}

func main() {
	var (
		n           int
		a, sum, avg float64
		all         []float64
	)
	_, _ = fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%f\n", &a)
		all = append(all, a)
	}
	sort.Float64s(all)
	for i := 1; i < n-1; i++ {
		sum += all[i]
	}
	avg = sum / float64(n-2)
	fmt.Printf("%.2f ", avg)
	maxF := -1.0
	for i := 1; i < n-1; i++ {
		t := abs2095(all[i] - avg)
		if t > maxF {
			maxF = t
		}
	}
	fmt.Printf("%.2f", maxF)
}
