package main

import "fmt"

//func firstBadVersion(n int) int {
//	if isBadVersion(1) {
//		return 1
//	}
//	return firstBadVersion2(1, n)
//}
//
////如果直接一个个查的话，表面看起来是一次循环，但是其实它里面那个东西
////调用起来极其耗时，所以只能二分查找了
//func firstBadVersion2(low, high int) int {
//	if low+1 == high {
//		return high
//	}
//	if isBadVersion((low + high) / 2) {
//		high = (low + high) / 2
//		return firstBadVersion2(low, high)
//	} else {
//		low = (low + high) / 2
//		return firstBadVersion2(low, high)
//	}
//}
//
//func isBadVersion(version int) bool {
//	if version < 194634574 {
//		return false
//	}
//	return true
//}

func firstBadVersion(n int) int {
	l, r, m := 1, n, 0
	for l < r {
		m = (l + r) / 2
		if isBadVersion(m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func isBadVersion(version int) bool {
	if version >= 2 {
		return true
	}
	return false
}

func main() {
	fmt.Println(firstBadVersion(3))
}
