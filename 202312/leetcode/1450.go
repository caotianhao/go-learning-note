package main

import "fmt"

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	cnt, l := 0, len(startTime)
	for i := 0; i < l; i++ {
		if queryTime >= startTime[i] && queryTime <= endTime[i] {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(busyStudent([]int{9, 8, 7, 6, 5, 4, 3, 2, 1},
		[]int{10, 10, 10, 10, 10, 10, 10, 10, 10}, 5))
}
