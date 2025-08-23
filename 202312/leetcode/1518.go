package main

import "fmt"

func numWaterBottles(numBottles int, numExchange int) int {
	water, empty := numBottles, numBottles
	water += empty / numExchange
	empty = empty - empty/numExchange*numExchange + empty/numExchange
	for empty >= numExchange {
		water += empty / numExchange
		empty = empty - empty/numExchange*numExchange + empty/numExchange
	}
	return water
}

func main() {
	fmt.Println(numWaterBottles(15, 4))
}
