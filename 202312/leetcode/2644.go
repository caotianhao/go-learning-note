package main

import (
	"fmt"
	"math"
)

func maxDivScore(nums, divisors []int) int {
	d, maxDivisorCnt, res := make([]int, 0), -1, math.MaxInt64
	for i := 0; i < len(divisors); i++ {
		cnt := 0
		for j := 0; j < len(nums); j++ {
			if nums[j]%divisors[i] == 0 {
				cnt++
			}
		}
		d = append(d, cnt)
	}
	for _, v := range d {
		if v > maxDivisorCnt {
			maxDivisorCnt = v
		}
	}
	for i, v := range divisors {
		if d[i] == maxDivisorCnt && v < res {
			res = v
		}
	}
	return res
}

func main() {
	fmt.Println(maxDivScore([]int{4, 7, 9, 3, 9}, []int{5, 2, 3}))
}
