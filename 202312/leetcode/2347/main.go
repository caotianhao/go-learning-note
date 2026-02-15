package main

import "fmt"

func bestHand(ranks []int, suits []byte) string {
	ranksMap, suitsMap := map[int]int{}, map[byte]int{}
	for _, v := range ranks {
		ranksMap[v]++
	}
	for _, v := range suits {
		suitsMap[v]++
	}
	if len(suitsMap) == 1 {
		return "Flush"
	}
	maxN := 0
	for _, v := range ranksMap {
		if v > maxN {
			maxN = v
		}
	}
	if maxN >= 3 {
		return "Three of a Kind"
	} else if maxN == 2 {
		return "Pair"
	} else {
		return "High Card"
	}
}

func main() {
	fmt.Println(bestHand([]int{1, 2, 3, 4, 5}, []byte{'a', 'a', 'a', 'a', 'a'}))
}
