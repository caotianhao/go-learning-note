package main

import "fmt"

func findTheWinner(n int, k int) int {
	player, begin := make([]int, n), 0
	for i := 1; i <= n; i++ {
		player[i-1] = i
	}
	for len(player) != 1 {
		t := (begin + k - 1) % len(player)
		player = remove1823(player, t)
		begin = t
	}
	return player[0]
}

func remove1823(arr []int, i int) (res []int) {
	for v := range arr {
		if v != i {
			res = append(res, arr[v])
		}
	}
	return
}

func main() {
	fmt.Println(findTheWinner(18, 9))
}
