package main

import (
	"fmt"
	"math"
)

func threeSumClosest(nums []int, target int) (r int) {
	l := len(nums)
	m := math.MaxInt64
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			for k := j + 1; k < l; k++ {
				t1 := nums[i] + nums[j] + nums[k]
				t2 := abs16(t1 - target)
				if t2 < m {
					m = t2
					r = t1
				}
			}
		}
	}
	return
}

func abs16(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))
	//arr := make([]int, 999)
	//for i := 0; i < 999; i++ {
	//	arr[i] = i + 63 - 96*14 + 415 - 999/3
	//}
	//fmt.Println(arr)
}
