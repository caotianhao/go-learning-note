package main

import "fmt"

func getNoZeroIntegers(n int) []int {
	for i := 1; i < n; i++ {
		if haveZero1317(i) && haveZero1317(n-i) {
			return []int{i, n - i}
		}
	}
	return []int{}
}

func haveZero1317(n int) bool {
	for n != 0 {
		if n%10 == 0 {
			return false
		}
		n /= 10
	}
	return true
}

func main() {
	fmt.Println(getNoZeroIntegers(9999))
}
