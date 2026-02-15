package main

import (
	"fmt"
	"time"
)

func nthUglyNumber2(n int) int {
	//cnt := 0
	//for i := 1; ; i += 1 {
	//	if isUgly264(i) {
	//		cnt++
	//	}
	//	if cnt == n {
	//		return i
	//	}
	//}

	uglyMap, cnt := map[int]int{}, 0
	for i := 1; cnt < n; i += 1 {
		if isUgly264(i) {
			uglyMap[cnt] = i
			cnt++
		}
	}
	return uglyMap[n-1]
}

//丑数就是只包含质因数 2、3 和/或 5 的正整数
//1 通常被视为丑数
func isUgly264(n int) bool {
	for n%2 == 0 || n%3 == 0 || n%5 == 0 {
		if n%2 == 0 {
			n /= 2
		}
		if n%3 == 0 {
			n /= 3
		}
		if n%5 == 0 {
			n /= 5
		}
	}
	if n == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	time0 := time.Now()
	ugly := nthUglyNumber2(1100)
	time1 := time.Since(time0)
	fmt.Println(ugly, time1)

	//fmt.Println(isUgly264(81))
	//time2 := time.Now()
	//for i := 0; i < 10000000; i++ {
	//	isUgly264(4000000000)
	//}
	//time3 := time.Since(time2)
	//fmt.Println(time3)
}
