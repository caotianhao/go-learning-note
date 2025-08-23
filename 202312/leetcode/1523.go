package main

import "fmt"

func countOdds(low int, high int) int {
	return (high-low)/2 + ((low % 2) | (high % 2))
}

func main() {
	fmt.Println(countOdds(18, 999))
}
