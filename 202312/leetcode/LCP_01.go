package main

import "fmt"

func game(guess []int, answer []int) int {
	cnt := 0
	for i := 0; i < 3; i++ {
		if guess[i] == answer[i] {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(game([]int{1, 2, 3}, []int{3, 2, 1}))
}
