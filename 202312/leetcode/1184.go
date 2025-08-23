package main

import (
	"fmt"
	"math"
)

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	clockwise, allDistance := 0, 0
	for _, v := range distance {
		allDistance += v
	}
	for i := start; i != destination; i = (i + 1) % len(distance) {
		clockwise += distance[i]
	}
	return int(math.Min(float64(clockwise), float64(allDistance-clockwise)))
}

func main() {
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 3))
	fmt.Println(distanceBetweenBusStops([]int{7, 10, 1, 12, 11, 14, 5, 0}, 7, 2))
}
