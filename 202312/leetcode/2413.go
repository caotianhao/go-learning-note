package main

import "fmt"

//给你一个正整数 n ，返回 2 和 n 的最小公倍数（正整数）。
//func smallestEvenMultiple(n int) int {
//	if n%2 == 1 {
//		return 2 * n
//	} else {
//		return n
//	}
//}

func smallestEvenMultiple(n int) int {
	if n&1 == 0 {
		return n
	}
	return n * 2
}

func main() {
	fmt.Println(smallestEvenMultiple(1))
}
