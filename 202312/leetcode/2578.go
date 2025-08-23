package main

import (
	"fmt"
	"math"
	"sort"
)

// 原来理解错题意了，以为只能从某处分割为左右两半
//func splitNum(num int) int {
//	str := strconv.FormatInt(int64(num), 10)
//	l, res := len(str), math.MaxInt64
//	for i := 1; i < l; i++ {
//		left, right := str[:i], str[i:]
//		a, _ := strconv.ParseInt(left, 10, 64)
//		b, _ := strconv.ParseInt(right, 10, 64)
//		fmt.Println(getMin2578(int(a)), getMin2578(int(b)))
//		res = min2578(res, getMin2578(int(a))+getMin2578(int(b)))
//	}
//	return res
//}
//
//func getMin2578(num int) (res int) {
//	arr := make([]int, 0)
//	for num != 0 {
//		arr = append(arr, num%10)
//		num /= 10
//	}
//	sort.Ints(arr)
//	l := len(arr)
//	for i := l - 1; i >= 0; i-- {
//		res += arr[i] * pow2578(10, l-i-1)
//	}
//	return res
//}
//
//func pow2578(a, b int) int {
//	return int(math.Pow(float64(a), float64(b)))
//}
//
//
//func min2578(a, b int) int {
//	if a > b {
//		return b
//	}
//	return a
//}

func splitNum(num int) int {
	arr := make([]int, 0)
	for num != 0 {
		arr = append(arr, num%10)
		num /= 10
	}
	sort.Ints(arr)
	l, a, b := len(arr), make([]int, 0), make([]int, 0)
	for i := 0; i < l; i++ {
		if i%2 == 0 {
			a = append(a, arr[i])
		} else {
			b = append(b, arr[i])
		}
	}
	aa, bb := 0, 0
	for i := len(a) - 1; i >= 0; i-- {
		aa += a[i] * pow2578(10, len(a)-i-1)
	}
	for i := len(b) - 1; i >= 0; i-- {
		bb += b[i] * pow2578(10, len(b)-i-1)
	}
	return aa + bb
}

func pow2578(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func main() {
	fmt.Println(splitNum(19))        // 10
	fmt.Println(splitNum(4325))      // 59
	fmt.Println(splitNum(687))       // 75
	fmt.Println(splitNum(10))        // 1
	fmt.Println(splitNum(671835844)) // 18046
	//fmt.Println(getMin2578(450000))
}
