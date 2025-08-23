package main

import (
	"fmt"
	"sort"
)

func findRelativeRanks(score []int) []string {
	trueSlice, ret := make([]int, 0), make([]string, 0)
	for _, v := range score {
		trueSlice = append(trueSlice, v)
	}
	sort.Slice(score, func(i, j int) bool {
		return score[i] > score[j]
	})
	rewardMap := map[int]string{}
	for i, v := range score {
		if i == 0 {
			rewardMap[v] = "Gold Medal"
		} else if i == 1 {
			rewardMap[v] = "Silver Medal"
		} else if i == 2 {
			rewardMap[v] = "Bronze Medal"
		} else {
			rewardMap[v] = fmt.Sprintf("%d", i+1)
		}
	}
	for _, v := range trueSlice {
		ret = append(ret, rewardMap[v])
	}
	return ret
}

func main() {
	fmt.Println(findRelativeRanks([]int{5, 4, 3, 2, 1}))
	fmt.Println(findRelativeRanks([]int{10, 3, 8, 9, 4}))
	fmt.Println(findRelativeRanks([]int{10, 9, 36}))
}
