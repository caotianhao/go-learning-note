package main

import (
	"fmt"
	"math"
	"sort"
)

func largestInteger(num int) int {
	numSlice, joSlice := make([]int, 0), make([]int, 0)
	j, o, ret := make([]int, 0), make([]int, 0), make([]int, 0)
	for num != 0 {
		numSlice = append(numSlice, num%10)
		num /= 10
	}
	numSlice = reverse2231(numSlice)
	for _, v := range numSlice {
		joSlice = append(joSlice, v%2)
		if v%2 == 0 {
			o = append(o, v)
		} else {
			j = append(j, v)
		}
	}
	sort.Ints(j)
	sort.Ints(o)
	reverse2231(j)
	reverse2231(o)
	jj, oo, want := 0, 0, 0
	for i := 0; jj != len(j) || oo != len(o); i++ {
		if joSlice[i] == 1 {
			ret = append(ret, j[jj])
			jj++
		} else {
			ret = append(ret, o[oo])
			oo++
		}
		if i == len(numSlice) {
			break
		}
	}
	ret = reverse2231(ret)
	for i := 0; i < len(ret); i++ {
		want += ret[i] * int(math.Pow(10.0, float64(i)))
	}
	return want
}

func reverse2231(numSlice []int) []int {
	l := len(numSlice)
	for i := 0; i < l/2; i++ {
		numSlice[i], numSlice[l-i-1] = numSlice[l-i-1], numSlice[i]
	}
	return numSlice
}

func main() {
	fmt.Println(largestInteger(288888))
}
