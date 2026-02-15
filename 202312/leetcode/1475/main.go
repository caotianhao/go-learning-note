package main

import "fmt"

func finalPrices(prices []int) []int {
	l, truePrice := len(prices), make([]int, 0)
	for i := 0; i < l-1; i++ {
		haveDiscount := false
		for j := i + 1; j < l; j++ {
			if prices[j] <= prices[i] {
				truePrice = append(truePrice, prices[i]-prices[j])
				haveDiscount = true
				break
			}
		}
		if !haveDiscount {
			truePrice = append(truePrice, prices[i])
		}
	}
	truePrice = append(truePrice, prices[l-1])
	return truePrice
}

func main() {
	fmt.Println(finalPrices([]int{8, 4, 6, 2, 3}))
	fmt.Println(finalPrices([]int{1, 2, 3, 4, 5}))
	fmt.Println(finalPrices([]int{10, 1, 1, 6, 8, 4, 6, 2, 3, 1, 1, 6, 8, 4, 6, 2, 3, 1, 1,
		6, 8, 4, 6, 2, 3, 1, 1, 6, 8, 4, 6, 2, 3, 1, 1, 6, 8, 4, 6, 2, 3, 1, 1, 6, 8, 4, 6,
		2, 3, 1, 1, 6, 8, 4, 6, 2, 3, 1, 1, 6, 8, 4, 6, 2, 3, 1, 1, 6, 8, 4, 6, 2, 3, 1, 1,
		6, 8, 4, 6, 2, 3}))
	fmt.Println(finalPrices([]int{1, 1, 1, 1, 1, 1, 1, 1, 1}))
}
