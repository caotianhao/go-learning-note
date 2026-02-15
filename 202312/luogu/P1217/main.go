package main

import (
	"fmt"
	"sort"
)

func getP1217() (p []int) {
	// 1,2
	p = append(p, 2, 3, 5, 7, 11)
	// 3,4
	for i := 1; i <= 9; i += 2 {
		for j := 0; j <= 9; j++ {
			a, b := 100*i+10*j+i, 1000*i+100*j+10*j+i
			if isPrime1217(a) {
				p = append(p, a)
			}
			if isPrime1217(b) {
				p = append(p, b)
			}
		}
	}
	// 5,6
	for i := 1; i <= 9; i += 2 {
		for j := 0; j <= 9; j++ {
			for k := 0; k <= 9; k++ {
				a, b := 10000*i+1000*j+100*k+10*j+i, 100000*i+10000*j+1000*k+100*k+10*j+i
				if isPrime1217(a) {
					p = append(p, a)
				}
				if isPrime1217(b) {
					p = append(p, b)
				}
			}
		}
	}
	// 7,8
	for i := 1; i <= 9; i += 2 {
		for j := 0; j <= 9; j++ {
			for k := 0; k <= 9; k++ {
				for m := 0; m <= 9; m++ {
					a := 1000000*i + 100000*j + 10000*k + 1000*m + 100*k + 10*j + i
					b := 10000000*i + 1000000*j + 100000*k + 10000*m + 1000*m + 100*k + 10*j + i
					if isPrime1217(a) {
						p = append(p, a)
					}
					if isPrime1217(b) {
						p = append(p, b)
					}
				}
			}
		}
	}
	return
}

func isPrime1217(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var a, b int
	_, _ = fmt.Scanf("%d %d", &a, &b)
	arr := getP1217()
	sort.Ints(arr)
	for _, v := range arr {
		if v >= a && v <= b {
			fmt.Println(v)
		}
	}
}
