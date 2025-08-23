package main

import (
	"fmt"
	"sort"
)

func halfQuestions(questions []int) int {
	sum := 0
	mapLCS02, sliceLCS02 := map[int]int{}, make([]int, 0)
	for _, v := range questions {
		mapLCS02[v]++
	}
	for _, v := range mapLCS02 {
		sliceLCS02 = append(sliceLCS02, v)
	}
	sort.Slice(sliceLCS02, func(i, j int) bool {
		return sliceLCS02[i] > sliceLCS02[j]
	})
	for i := 0; i < len(sliceLCS02); i++ {
		sum += sliceLCS02[i]
		if sum >= len(questions)/2 {
			return i + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(halfQuestions([]int{1, 5, 1, 3, 4, 5, 2, 5, 3, 3, 8, 6})) //2
	fmt.Println(halfQuestions([]int{2, 1, 6, 2}))                         //1
}
