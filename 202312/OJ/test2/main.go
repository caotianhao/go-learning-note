package main

import "fmt"

func t2(nms []int) []int {
	n := len(nms)
	res := make([]int, n)
	res[0] = nms[0]
	for i := 1; i < n; i++ {
		res[i] = res[i-1] + nms[i]
	}
	return res
}

func t22(nums []int) []int {
	n := len(nums)
	p := make([]int, n+1)
	for i := 1; i <= n; i++ {
		p[i] = p[i-1] + nums[i-1]
	}
	return p
}

func main() {
	a := make([]int, 0)
	for i := 1; i < 12; i++ {
		a = append(a, i*i*i+9*i+4*i*i)
	}
	fmt.Println(a)
	fmt.Println(t2(a))
	fmt.Println(t22(a))
}
