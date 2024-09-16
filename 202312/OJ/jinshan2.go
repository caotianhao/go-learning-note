package main

import (
	"fmt"
	"math"
)

func maxFruitValue(apple, peach, fruitA, fruitB int) int {
	// apple + apple + peach = fruitA
	// apple + peach + peach = fruitB
	value := math.MinInt64
	maxA, maxB := min(apple/2, peach), min(peach/2, apple)
	for i := 0; i <= maxA; i++ {
		value = max(value, i*fruitA+min((peach-i)/2, apple-2*i)*fruitB)
	}
	for i := 0; i <= maxB; i++ {
		value = max(value, i*fruitB+min((apple-i)/2, peach-2*i)*fruitA)
	}
	if value == math.MinInt64 {
		return 0
	}
	return value
}

func main() {
	apple, peach, fruitA, fruitB := 0, 0, 0, 0
	_, _ = fmt.Scanf("%d %d %d %d", &apple, &peach, &fruitA, &fruitB)
	fmt.Println(maxFruitValue(apple, peach, fruitA, fruitB))
}
