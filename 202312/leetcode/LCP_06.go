package main

import "fmt"

func minCount(coins []int) (cnt int) {
	for _, v := range coins {
		if v%2 == 0 {
			cnt += v / 2
		} else {
			cnt += v/2 + 1
		}
	}
	return
}

func main() {
	fmt.Println(minCount([]int{4, 2, 1}))  //4
	fmt.Println(minCount([]int{2, 3, 10})) //8
}
