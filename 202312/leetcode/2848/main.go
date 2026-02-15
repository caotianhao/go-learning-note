package main

import "fmt"

func numberOfPoints(nums [][]int) int {
	hm := make(map[int]struct{})
	for _, v := range nums {
		for i := v[0]; i <= v[1]; i++ {
			hm[i] = struct{}{}
		}
	}
	return len(hm)
}

func main() {
	fmt.Println(numberOfPoints([][]int{{3, 6}, {1, 5}, {4, 7}}))
	fmt.Println(numberOfPoints([][]int{{1, 3}, {5, 8}}))
}
