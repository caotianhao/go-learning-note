package main

import "fmt"

func findNonMinOrMax(nums []int) int {
	if len(nums) < 3 {
		return -1
	} else {
		minN, maxN := 101, 0
		for _, v := range nums {
			if v > maxN {
				maxN = v
			}
			if v < minN {
				minN = v
			}
		}
		for _, v := range nums {
			if v != minN && v != maxN {
				return v
			}
		}
		return -1
	}
}

func main() {
	fmt.Println(findNonMinOrMax([]int{1, 2, 3}))
}
