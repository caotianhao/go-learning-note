package main

import (
	"fmt"
	"math"
)

// cout 和 go 的输出不一样
// 洛谷没有官方题解，更对 go 不友好...
func main() {
	var n int
	const p = 3.141593
	_, _ = fmt.Scanf("%d", &n)
	switch n {
	case 1:
		fmt.Printf("I love Luogu!")
	case 2:
		fmt.Printf("%d %d", 6, 4)
	case 3:
		fmt.Printf("%d\n%d\n%d", 3, 12, 2)
	case 4:
		fmt.Printf("%.3f", 500.0/3.0)
	case 5:
		fmt.Printf("%d", 15)
	case 6:
		fmt.Printf("%.4f", math.Sqrt(117.0))
	case 7:
		fmt.Printf("%d\n%d\n%d", 110, 90, 0)
	case 8:
		fmt.Printf("%.4f\n%.4f\n%.3f", 10.0*p, 25.0*p, 500.0*p/3.0)
	case 9:
		fmt.Printf("%d", 22)
	case 10:
		fmt.Printf("%d", 9)
	case 11:
		fmt.Printf("%.4f", 100.0/3.0)
	case 12:
		fmt.Printf("%d\n%c", 13, 'R')
	case 13:
		fmt.Printf("%d", int(math.Floor(math.Pow(1064.0*4.0*p/3.0, 1.0/3.0))))
	case 14:
		fmt.Printf("%d", 50)
	}
}
