package main

import "fmt"

func numPrimeArrangements(n int) int {
	isPrime := func(num int) bool {
		if num == 1 {
			return false
		}
		for i := 2; i < num; i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}
	prime, ans := 0, 1
	for i := 1; i <= n; i++ {
		if isPrime(i) {
			prime++
		}
	}
	for i := 1; i <= prime; i++ {
		ans *= i
		for ans > 1000000007 {
			ans -= 1000000007
		}
	}
	for i := 1; i <= n-prime; i++ {
		ans *= i
		for ans > 1000000007 {
			ans -= 1000000007
		}
	}
	return ans
}

func main() {
	//fmt.Println(numPrimeArrangements(1))
	//fmt.Println(numPrimeArrangements(2))
	//fmt.Println(numPrimeArrangements(3))
	//fmt.Println(numPrimeArrangements(4))
	fmt.Println(numPrimeArrangements(5))
}
