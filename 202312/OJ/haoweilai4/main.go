package main

import "fmt"

func countPrime(n int) int {
	cnt := 0
	isNotPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		if !isNotPrime[i] {
			cnt++
			for j := i * i; j < n; j += i {
				isNotPrime[j] = true
			}
		}
	}
	return cnt
}

func main() {
	n := 0
	_, _ = fmt.Scanf("%d", &n)
	fmt.Printf("%d", countPrime(n))
}
