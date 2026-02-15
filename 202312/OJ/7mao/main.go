package main

import (
	"fmt"
	"strconv"
)

func getSum(ss string) (s int) {
	for _, v := range ss {
		s += int(v - '0')
	}
	return
}

func gcd(a, b int) int {
	if b != 0 {
		return gcd(b, a%b)
	}
	return a
}

func main() {
	n := 0
	for {
		f, _ := fmt.Scanf("%d", &n)
		if f == 0 {
			break
		}
		sss := 0
		for i := 2; i < n; i++ {
			q := strconv.FormatInt(int64(n), i)
			sss += getSum(q)
		}
		fmt.Print(sss / gcd(sss, n-2))
		fmt.Print("/")
		fmt.Print((n - 2) / gcd(sss, n-2))
	}
}
