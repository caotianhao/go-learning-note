package main

import "fmt"

//func mySqrt(x int) int {
//	myReturn := 0
//	if x == 1 || x == 2 || x == 3 {
//		return 1
//	}
//	if x == 4 || x == 5 || x == 6 || x == 7 || x == 8 {
//		return 2
//	}
//	for i := 1; i < x/2; i++ {
//		if i*i == x {
//			return i
//		}
//		if (i+1)*(i+1) > x {
//			myReturn = i
//			break
//		}
//	}
//	return myReturn
//}

func mySqrt(x int) int {
	left, right := 0, x
	for left < right {
		mid := (left + right + 1) >> 1
		if mid*mid > x {
			right = mid - 1
		}
		if mid*mid < x {
			left = mid
		}
		if mid*mid == x {
			return mid
		}
	}
	return left
}

func main() {
	fmt.Println(mySqrt(10))
	//fmt.Println(mySqrt(1))
	//fmt.Println(mySqrt(0))
	for i := 0; i <= 9; i++ {
		fmt.Println(mySqrt(i))
	}
}
