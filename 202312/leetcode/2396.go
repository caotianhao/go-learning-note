package main

import (
	"fmt"
	"reflect"
)

func isStrictlyPalindromic(n int) bool {
	for i := 2; i <= n-2; i++ {
		if !reflect.DeepEqual(jinZhi(n, i), reverse2396(jinZhi(n, i))) {
			return false
		}
	}
	return true
}

func jinZhi(num, j int) (ret []int) {
	for num != 0 {
		ret = append(ret, num%j)
		num /= j
	}
	return
}

func reverse2396(arr []int) []int {
	l := len(arr)
	for i := 0; i < l>>1; i++ {
		arr[i], arr[l-i-1] = arr[l-i-1], arr[i]
	}
	return arr
}

func main() {
	//fmt.Println(reverse2396([]int{1, 2, 3, 4}))
	//fmt.Println(jinZhi(10, 2))
	fmt.Println(isStrictlyPalindromic(9))
	fmt.Println(isStrictlyPalindromic(4))
	//for i := 4; i <= 100000; i++ {
	//	if isStrictlyPalindromic(i) {
	//		fmt.Println(i)
	//	}
	//}
}
