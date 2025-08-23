package main

import "fmt"

func averageValue(nums []int) int {
	cnt, sum := 0, 0
	for _, v := range nums {
		if v&1 == 0 && v%3 == 0 {
			sum += v
			cnt++
		}
	}
	if cnt == 0 {
		return cnt
	}
	return sum / cnt
}

func main() {
	fmt.Println(averageValue([]int{1, 3, 6, 10, 12, 15}))
	fmt.Println(averageValue([]int{1, 2, 4, 10, 7}))
}
