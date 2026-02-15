package main

import (
	"fmt"
	"math/bits"
)

func countPrimeSetBits(left int, right int) int {
	cnt := 0
	for i := left; i <= right; i++ {
		if isPrime762(bits.OnesCount(uint(i))) {
			cnt++
		}
	}
	return cnt
}

func isPrime762(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(countPrimeSetBits(6, 10))
	fmt.Println(countPrimeSetBits(6, 6))
}
