package main

import (
	"fmt"
	"math"
	"sort"
)

func minMovesToSeat(seats []int, students []int) int {
	l, minMove := len(seats), 0
	sort.Ints(seats)
	sort.Ints(students)
	for i := 0; i < l; i++ {
		minMove += int(math.Abs(float64(seats[i] - students[i])))
	}
	return minMove
}

func main() {
	fmt.Println(minMovesToSeat([]int{4, 1, 5, 9}, []int{1, 3, 2, 6}))
}
