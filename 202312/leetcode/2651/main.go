package main

import "fmt"

func findDelayedArrivalTime(arrivalTime int, delayedTime int) int {
	return (arrivalTime + delayedTime) % 24
}

func main() {
	fmt.Println(findDelayedArrivalTime(18, 20))
}
