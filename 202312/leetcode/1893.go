package main

import "fmt"

func isCovered1(ranges [][]int, left int, right int) bool {
	for i := left; i <= right; i++ {
		flag := false
		for _, v := range ranges {
			if i >= v[0] && i <= v[1] {
				flag = true
				break
			}
		}
		if !flag {
			return false
		}
	}
	return true
}

// 差分数组
func isCovered(ranges [][]int, left, right int) bool {
	diff := make([]int, 52)
	for _, v := range ranges {
		diff[v[0]]++
		diff[v[1]+1]--
	}
	prefixSum := make([]int, 52)
	for i := 1; i < 52; i++ {
		prefixSum[i] += diff[i] + prefixSum[i-1]
	}
	for i := left; i <= right; i++ {
		if prefixSum[i] <= 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isCovered1([][]int{{1, 2}, {2, 3}}, 2, 3))
	fmt.Println(isCovered([][]int{{1, 2}, {2, 3}}, 2, 3))
}
