package main

import (
	"fmt"
	"gogogo/project02"
	"time"
)

//func isPrime(n int) bool {
//	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
//		if n%i == 0 {
//			return false
//		}
//	}
//	return true
//}
//
//func printPrime(num1 int, num2 int) {
//	for i := num1; i <= num2; i++ {
//		if isPrime(i) {
//			//fmt.Println(i)
//		}
//	}
//}

func main() {
	var a, b int
	_, _ = fmt.Scanf("%d", &a)
	_, _ = fmt.Scanf("%d", &b)
	t0 := time.Now()
	project02.PrintPrime(a, b)
	t := time.Since(t0)
	fmt.Println(t)
}
