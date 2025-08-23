package main

import (
	"fmt"
	"math"
)

func maxDistance(colors []int) int {
	l, maxN := len(colors), -1
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if colors[i] != colors[j] && int(math.Abs(float64(i-j))) > maxN {
				maxN = int(math.Abs(float64(i - j)))
			}
		}
	}
	return maxN
}

func main() {
	fmt.Println(maxDistance([]int{1, 1, 1, 6, 1, 1, 1}))
}
