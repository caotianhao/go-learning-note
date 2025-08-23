package main

import "fmt"

//啥原理啊...
func minCostToMoveChips(position []int) int {
	l, cnt, cnt1, cnt0 := len(position), 0, 0, 0
	for i := 0; i < l; i++ {
		if position[i]%2 == 0 {
			cnt0++
		} else {
			cnt1++
		}
		if cnt0 > cnt1 {
			cnt = cnt1
		} else {
			cnt = cnt0
		}
	}
	return cnt
}

func main() {
	fmt.Println(minCostToMoveChips([]int{2, 2, 2, 3, 3}))
}
