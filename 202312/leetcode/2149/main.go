package main

import "fmt"

func rearrangeArray(nums []int) (res []int) {
	pos, neg := make([]int, 0), make([]int, 0)
	for _, v := range nums {
		if v > 0 {
			pos = append(pos, v)
		} else {
			neg = append(neg, v)
		}
	}
	for c := 0; c < len(nums); {
		if c&1 == 0 {
			res = append(res, pos[c/2])
		} else {
			res = append(res, neg[c/2])
		}
		c++
	}
	return
}

func main() {
	fmt.Println(rearrangeArray([]int{3, 1, -2, -5, 2, -4}))
	fmt.Println(rearrangeArray([]int{-1, 1}))
}
