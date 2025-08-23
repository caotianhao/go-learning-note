package main

import "fmt"

func distanceTraveled(mainTank int, additionalTank int) int {
	dis := 0
	for mainTank >= 5 {
		mainTank -= 5
		dis += 50
		if additionalTank > 0 {
			additionalTank--
			mainTank++
		}
	}
	return dis + 10*mainTank
}

func main() {
	fmt.Println(distanceTraveled(5, 10))
}
