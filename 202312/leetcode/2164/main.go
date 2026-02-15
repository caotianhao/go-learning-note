package main

import (
	"fmt"
	"sort"
)

func sortEvenOdd(nums []int) []int {
	l, ji, ou, ret := len(nums), make([]int, 0), make([]int, 0), make([]int, 0)
	if l < 3 {
		return nums
	}
	for i := 0; i < l; i += 2 {
		ou = append(ou, nums[i])
	}
	for i := 1; i < l; i += 2 {
		ji = append(ji, nums[i])
	}
	sort.Ints(ou)
	sort.Slice(ji, func(i, j int) bool {
		return ji[i] > ji[j]
	})
	flag := 0
	for i := 0; i < l/2; i++ {
		if flag == 0 {
			ret = append(ret, ou[i])
			flag = 1 - flag
		}
		if flag == 1 {
			ret = append(ret, ji[i])
			flag = 1 - flag
		}
	}
	if l%2 == 1 {
		ret = append(ret, ou[len(ou)-1])
	}
	return ret
}

func main() {
	fmt.Println(sortEvenOdd([]int{4, 1, 2, 3, 6}))
	//fmt.Println(sortEvenOdd([]int{0}))
	//fmt.Println(sortEvenOdd([]int{1}))
	//fmt.Println(sortEvenOdd([]int{7}))
	//fmt.Println(sortEvenOdd([]int{1, 0}))
	//fmt.Println(sortEvenOdd([]int{7, 11}))
}
