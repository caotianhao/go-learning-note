package main

import "fmt"

func distributeCandies(candies, numPeople int) []int {
	ans := make([]int, numPeople)
	for i := 1; ; i++ {
		if candies-i >= 0 {
			ans[(i-1)%numPeople] += i
			candies -= i
		} else {
			ans[(i-1)%numPeople] += candies
			break
		}
	}
	return ans
}

func main() {
	fmt.Println(distributeCandies(7, 4))
	fmt.Println(distributeCandies(10, 3))
	fmt.Println(distributeCandies(2993, 15))
}
