package main

import (
	"fmt"
	"math"
)

func distributeCandies575(candyType []int) int {
	candyMap := map[int]int{}
	for _, v := range candyType {
		candyMap[v]++
	}
	return int(math.Min(float64(len(candyMap)), float64(len(candyType)/2)))
}

func main() {
	fmt.Println(distributeCandies575([]int{1, 1, 2, 2, 3, 3}))
}
