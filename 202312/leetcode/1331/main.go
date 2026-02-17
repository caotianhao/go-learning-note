package main

import (
	"fmt"
	"sort"
)

func arrayRankTransform(arr []int) []int {
	var tmpArr, ret []int
	map1331, ind := map[int]int{}, 1
	for _, v := range arr {
		tmpArr = append(tmpArr, v)
	}
	sort.Ints(arr)
	for _, v := range arr {
		if _, ok := map1331[v]; ok {
			continue
		}
		map1331[v] = ind
		ind++
	}
	for i := range tmpArr {
		ret = append(ret, map1331[tmpArr[i]])
	}
	return ret
}

func main() {
	fmt.Println(arrayRankTransform([]int{37, 12, 28, 9, 100, 56, 80, 5, 12}))
	fmt.Println(arrayRankTransform([]int{37}))
	fmt.Println(arrayRankTransform([]int{}))
}
