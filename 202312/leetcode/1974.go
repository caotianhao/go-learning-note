package main

import (
	"fmt"
	"math"
)

func minTimeToType(word string) int {
	l, slice1974, cnt := len(word), make([]int, 1), 0
	for i := 0; i < l; i++ {
		slice1974 = append(slice1974, int(word[i]-'a'))
	}
	for i := 0; i < len(slice1974)-1; i++ {
		cnt += func1974(slice1974[i], slice1974[i+1])
	}
	return cnt + l
}

func func1974(a, b int) int {
	m := int(math.Abs(float64(a - b)))
	n := 26 - m
	return int(math.Min(float64(m), float64(n)))
}

func main() {
	fmt.Println(minTimeToType("zjpc"))
}
