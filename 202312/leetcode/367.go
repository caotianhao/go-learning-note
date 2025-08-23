package main

import "fmt"

func isPerfectSquare(num int) bool {
	l, r := 0, num
	for l < r {
		mid := (l + r + 1) >> 1
		if mid*mid > num {
			r = mid - 1
		}
		if mid*mid < num {
			l = mid + 1
		}
		if mid*mid == num {
			return true
		}
	}
	if l*l == num {
		return true
	}
	return false
}

func main() {
	//fmt.Println(isPerfectSquare(18))
	//fmt.Println(isPerfectSquare(9))
	//fmt.Println(isPerfectSquare(1))
	//fmt.Println(isPerfectSquare(144))
	//fmt.Println(isPerfectSquare(2344))
	fmt.Println(isPerfectSquare(4))
}
