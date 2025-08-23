package main

import "fmt"

func buyChoco(prices []int, money int) int {
	min0, min1 := 101, 101
	for _, v := range prices {
		if v < min0 {
			min0, min1 = v, min0
		} else if v < min1 {
			min1 = v
		}
	}
	//fmt.Println(min0, min1)
	if min0+min1 <= money {
		return money - min1 - min0
	}
	return money
}

func main() {
	fmt.Println(buyChoco([]int{12, 4, 5, 6, 7, 1, 1, 4}, 3))
}
