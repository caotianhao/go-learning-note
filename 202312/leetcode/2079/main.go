package main

import "fmt"

func wateringPlants(plants []int, capacity int) (res int) {
	l, watered, ca := len(plants), 0, capacity
	for watered != l {
		ca -= plants[watered]
		if ca >= 0 {
			res, watered = res+1, watered+1
		} else {
			ca, res = capacity, res+2*watered
		}
	}
	return
}

func main() {
	fmt.Println(wateringPlants([]int{2, 2, 3, 3, 4, 2, 1, 3, 4, 2, 5, 1, 3}, 6))
}
