package main

import (
	"fmt"
	"math"
)

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	sumA, sumB, mapA, mapB := 0, 0, map[int]int{}, map[int]int{}
	for _, v := range aliceSizes {
		sumA += v
		mapA[v] = v
	}
	for _, v := range bobSizes {
		sumB += v
		mapB[v] = v
	}
	need := int(math.Abs(float64(sumA-sumB))) / 2
	for i := 0; i < len(aliceSizes); i++ {
		if _, ok := mapB[need+aliceSizes[i]]; ok && sumA+need == sumB-need {
			return []int{aliceSizes[i], need + aliceSizes[i]}
		}
		if _, ok := mapB[aliceSizes[i]-need]; ok && sumA-need == sumB+need {
			return []int{aliceSizes[i], aliceSizes[i] - need}
		}
	}
	return nil
}

func main() {
	fmt.Println(fairCandySwap([]int{1, 2, 5}, []int{2, 4}))
}
