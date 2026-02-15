package main

import "fmt"

func isWinner(player1, player2 []int) int {
	l := len(player1)
	s1, s2 := player1[0], player2[0]
	for i := 1; i < l; i++ {
		if player1[i-1] == 10 || (i-2 >= 0 && player1[i-2] == 10) {
			s1 += 2 * player1[i]
		} else {
			s1 += player1[i]
		}
	}
	for i := 1; i < l; i++ {
		if player2[i-1] == 10 || (i-2 >= 0 && player2[i-2] == 10) {
			s2 += 2 * player2[i]
		} else {
			s2 += player2[i]
		}
	}
	if s1 == s2 {
		return 0
	} else if s1 > s2 {
		return 1
	}
	return 2
}

func main() {
	fmt.Println(isWinner([]int{4, 10, 7, 9}, []int{6, 5, 2, 3}))
	fmt.Println(isWinner([]int{5, 6, 1, 10}, []int{5, 1, 10, 5}))
	fmt.Println(isWinner([]int{2, 3}, []int{4, 1}))
}
