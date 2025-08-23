package main

import (
	"fmt"
	"math"
	"sort"
)

func shortestToChar(s string, c byte) []int {
	l, index := len(s), make([]int, 0)
	slice821 := make([]int, l)
	for i := 0; i < l; i++ {
		if s[i] == c {
			index = append(index, i)
		}
	}
	for i := 0; i < len(index); i++ {
		slice821[index[i]] = 0
	}
	for i := 0; i < l; i++ {
		temp := make([]int, 0)
		for j := 0; j < len(index); j++ {
			temp = append(temp, int(math.Abs(float64(i-index[j]))))
		}
		sort.Ints(temp)
		slice821[i] = temp[0]
	}
	//fmt.Println(index)
	return slice821
}

func main() {
	fmt.Println(shortestToChar("loveleetcode", 'e'))
}
