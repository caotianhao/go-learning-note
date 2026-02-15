package main

import "fmt"

func stoneGame(piles []int) bool {
	return len(piles) > 1
}

func main() {
	fmt.Println(stoneGame([]int{5, 3, 4, 5}))
}
