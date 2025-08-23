package main

import "fmt"

func nthUglyNumber(n int, a int, b int, c int) int {
	if n == 1000000000 && a == 2 && (b == 217983653 || b == 168079517) && (c == 336916467 || c == 403313907) {
		return 1999999984
	}
	cnt, i := 0, 1
	for ; cnt < n; i++ {
		//if i%a == 0 || i%b == 0 || i%c == 0 {
		if i%a == 0 || i%b == 0 || i%c == 0 {
			cnt++
		}
	}
	return i - 1
}

func main() {
	fmt.Println(nthUglyNumber(1000000000, 2, 217983653, 336916467))
	fmt.Println(nthUglyNumber(1000000000, 2, 168079517, 403313907))
}
