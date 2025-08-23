package main

import "fmt"

func countPrimes(n int) (cnt int) {
	isNotPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		if !isNotPrime[i] {
			cnt++
			for j := i * i; j < n; j += i {
				isNotPrime[j] = true
			}
		}
	}
	return
}

func main() {
	fmt.Println(countPrimes(5000000))
	fmt.Println(countPrimes(0))
	fmt.Println(countPrimes(1))
	fmt.Println(countPrimes(2))
	fmt.Println(countPrimes(3))
}
