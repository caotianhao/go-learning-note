package main

import (
	"fmt"
	"sort"
)

func uniqueOccurrences(arr []int) bool {
	occurrences, tmp := map[int]int{}, make([]int, 0)
	for _, v := range arr {
		occurrences[v]++
	}
	for _, v := range occurrences {
		tmp = append(tmp, v)
	}
	sort.Ints(tmp)
	for i := 0; i < len(tmp)-1; i++ {
		if tmp[i] == tmp[i+1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(uniqueOccurrences([]int{5, 2, 2, 4, 4, 4, 6, 6, 6}))
}
