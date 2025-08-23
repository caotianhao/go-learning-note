package main

import "fmt"

func countTriples(n int) int {
	cnt := 0
	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			for k := 1; k <= n; k++ {
				if i*i+j*j == k*k {
					fmt.Println(i, j, k)
					cnt++
				}
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(countTriples(10))
}
