package main

import (
	"fmt"
	"time"
)

func makeIntegerBeautiful(n int64, target int) int64 {
	if n < 100 {
		tmp := everySum2457(n)
		var cnt int64
		for tmp > int64(target) {
			cnt++
			n += 1
			tmp = everySum2457(n)
		}
		return cnt
	} else {
		var i int64
		if everySum2457(n) <= int64(target) {
			return 0
		}
		for i = 10; ; i *= 10 {
			if everySum2457((n/i+1)*i) <= int64(target) {
				return (n/i+1)*i - n
			}
		}
	}
}

func everySum2457(n int64) (sum int64) {
	slice2457 := make([]int64, 0)
	for n != 0 {
		slice2457 = append(slice2457, n%int64(10))
		n /= int64(10)
	}
	for _, v := range slice2457 {
		sum += v
	}
	return
}

func main() {
	fmt.Println(makeIntegerBeautiful(16, 6))
	fmt.Println(makeIntegerBeautiful(467, 6))
	fmt.Println(makeIntegerBeautiful(1, 1))
	t := time.Now()
	k := makeIntegerBeautiful(10000000000000001, 1)
	t1 := time.Since(t)
	fmt.Println(k, t1)
}
