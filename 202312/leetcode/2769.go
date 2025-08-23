package main

import "fmt"

func theMaximumAchievableX(num int, t int) int {
	return num + t + t
}

func main() {
	fmt.Println(theMaximumAchievableX(5, 9))
}
