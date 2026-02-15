package main

import "fmt"

func findSpecialInteger(arr []int) int {
	hm := map[int]int{}
	for _, v := range arr {
		hm[v]++
	}
	for i, v := range hm {
		if v > len(arr)/4 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(findSpecialInteger([]int{1, 2, 3, 4, 4, 4, 6, 18, 5}))
}
