package main

import "fmt"

func simplifiedFractions(n int) []string {
	mySlice := make([]string, 0)
	for i := 2; i <= n; i++ {
		for j := 1; j <= n-1; j++ {
			if isTwoSumPrime(i, j) && j < i {
				mySlice = append(mySlice, fmt.Sprintf("%d", j)+"/"+fmt.Sprintf("%d", i))
			}
		}
	}
	return mySlice
}

func isTwoSumPrime(i int, j int) bool {
	if gcd2(i, j) == 1 {
		return true
	} else {
		return false
	}
}

func gcd2(i int, j int) int {
	if i < j {
		i, j = j, i
	}
	for i%j != 0 {
		return gcd2(j, i%j)
	}
	return j
}

func main() {
	//fmt.Println(gcd(6, 8))
	fmt.Println(simplifiedFractions(4))
}
