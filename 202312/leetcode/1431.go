package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	slice1431, l := make([]bool, 0), len(candies)
	maxN := 0
	for i := 0; i < l; i++ {
		if candies[i] > maxN {
			maxN = candies[i]
		}
	}
	for i := 0; i < l; i++ {
		if candies[i]+extraCandies >= maxN {
			slice1431 = append(slice1431, true)
		} else {
			slice1431 = append(slice1431, false)
		}
	}
	return slice1431
}

func main() {
	fmt.Println(kidsWithCandies([]int{12, 1, 12}, 10))
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
}
