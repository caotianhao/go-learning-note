package main

import (
	"fmt"
	"sort"
)

func numMovesStones(a int, b int, c int) []int {
	minMove, maxMove := -1, -1
	dis := append([]int{}, a, b, c)
	sort.Ints(dis)
	left, mid, right := dis[0], dis[1], dis[2]
	maxMove = right - left - 2
	if right-mid == 1 && mid-left == 1 {
		minMove = 0
	} else if (right-mid != 1 && mid-left == 1) ||
		(right-mid == 1 && mid-left != 1) ||
		(right-mid == 2) || (mid-left == 2) {
		minMove = 1
	} else {
		minMove = 2
	}
	return []int{minMove, maxMove}
}

func main() {
	fmt.Println(numMovesStones(1, 2, 5)) // 1 2
	fmt.Println(numMovesStones(4, 3, 2)) // 0 0
	fmt.Println(numMovesStones(3, 5, 1)) // 1,2
	fmt.Println(numMovesStones(1, 8, 9))
}
