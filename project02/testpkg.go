package project02

import "math"

func isPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func PrintPrime(num1 int, num2 int) {
	for i := num1; i <= num2; i++ {
		if isPrime(i) {
			//fmt.Println(i)
		}
	}
}
