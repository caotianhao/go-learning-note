package main

import (
	"fmt"
)

func jds1GoodMod(nums []int) (res []int) {
	n := len(nums)
	num := 1
	for i := 0; i < n; i++ {
		for num < 99999 {
			if (nums[i]+num)%(i+1) == 0 {
				res = append(res, num)
				num++
				break
			} else {
				num++
			}
		}
	}
	return
}

func main() {
	n := 0
	_, _ = fmt.Scanf("%d", &n)
	a := 0
	arr := make([]int, 0)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &a)
		arr = append(arr, a)
	}
	ans := jds1GoodMod(arr)
	for _, v := range ans {
		fmt.Printf("%d ", v)
	}
}
