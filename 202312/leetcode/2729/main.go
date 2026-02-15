package main

import (
	"fmt"
	"sort"
)

func isFascinating(n int) bool {
	if n > 329 || n%10 == 0 || n < 123 {
		return false
	}
	k := n*1000000 + n*2000 + n*3
	//fmt.Println(k)
	arr := make([]int, 0)
	for k != 0 {
		arr = append(arr, k%10)
		k /= 10
	}
	sort.Ints(arr)
	for i, v := range arr {
		if i+1 != v {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isFascinating(192))
}
