package main

import (
	"fmt"
	"math"
)

func reverseInt(n int) int {
	ans := 0
	for n != 0 {
		if ans < math.MinInt32/10 || ans > math.MaxInt32/10 {
			return 0
		}
		digits := n % 10
		n /= 10
		ans = ans*10 + digits
	}
	return ans
}

func main() {
	n := 0
	_, _ = fmt.Scanf("%d", &n)
	fmt.Printf("%d", reverseInt(n))
}
