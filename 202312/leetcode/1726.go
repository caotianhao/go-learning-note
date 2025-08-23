package main

import "fmt"

func tupleSameProduct(nums []int) (ans int) {
	n, hm := len(nums), make(map[int]int)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			hm[nums[i]*nums[j]]++
		}
	}
	for _, v := range hm {
		ans += v * (v - 1) * 4
	}
	return
}

func main() {
	fmt.Println(tupleSameProduct([]int{2, 3, 4, 6}))
}
