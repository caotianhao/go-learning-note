package main

import (
	"fmt"
	"sort"
)

func maxCoins(piles []int) (max int) {
	sort.Ints(piles)
	l := len(piles)
	for i, tmp := 0, 2; i < l/3; i, tmp = i+1, tmp+2 {
		max += piles[l-tmp]
	}
	return max
}

func main() {
	fmt.Println(maxCoins([]int{4, 6, 1, 2, 7, 2, 6, 9, 1, 75, 96, 11}))
	fmt.Println(maxCoins([]int{2, 4, 1, 2, 7, 8}))          //9
	fmt.Println(maxCoins([]int{2, 4, 5}))                   //4
	fmt.Println(maxCoins([]int{9, 8, 7, 6, 5, 1, 2, 3, 4})) //18
}
