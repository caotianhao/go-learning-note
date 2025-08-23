package main

import "fmt"

func findWinners(matches [][]int) [][]int {
	win, lose := map[int]int{}, map[int]int{}
	ret := make([][]int, 0)
	winner, loser := make([]int, 0), make([]int, 0)
	for _, v := range matches {
		win[v[0]]++
		lose[v[1]]++
	}
	for i := 1; i <= 100000; i++ {
		_, ok1 := win[i]
		q, ok2 := lose[i]
		if ok1 && !ok2 {
			winner = append(winner, i)
		}
		if (ok1 || ok2) && q == 1 {
			loser = append(loser, i)
		}
	}
	ret = append(ret, winner, loser)
	return ret
}

func main() {
	fmt.Println(findWinners([][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5},
		{4, 8}, {4, 9}, {10, 4}, {10, 9}}))
	fmt.Println(findWinners([][]int{{1, 3}}))
}
