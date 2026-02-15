package main

import "fmt"

func expectNumber(scores []int) int {
	mapLcp11 := map[int]int{}
	for _, v := range scores {
		mapLcp11[v]++
	}
	return len(mapLcp11)
}

func main() {
	fmt.Println(expectNumber([]int{1, 2, 3})) //3
	fmt.Println(expectNumber([]int{1, 1}))    //1
	fmt.Println(expectNumber([]int{1, 1, 2})) //2
}
