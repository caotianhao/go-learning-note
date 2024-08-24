package main

import (
	"fmt"
	"math"
)

func maxFruitValue(apple, peach, fruitA, fruitB int) int {
	// apple + apple + peach = fruitA
	// apple + peach + peach = fruitB
	value := math.MinInt64
	maxA, maxB := minFruit(apple/2, peach), minFruit(peach/2, apple)
	for i := 0; i <= maxA; i++ {
		value = maxValue(value, i*fruitA+minFruit((peach-i)/2, apple-2*i)*fruitB)
	}
	for i := 0; i <= maxB; i++ {
		value = maxValue(value, i*fruitB+minFruit((apple-i)/2, peach-2*i)*fruitA)
	}
	if value == math.MinInt64 {
		return 0
	}
	return value
}

func minFruit(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	apple, peach, fruitA, fruitB := 0, 0, 0, 0
	_, _ = fmt.Scanf("%d %d %d %d", &apple, &peach, &fruitA, &fruitB)
	fmt.Println(maxFruitValue(apple, peach, fruitA, fruitB))
}
