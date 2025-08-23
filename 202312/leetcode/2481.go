package main

import "fmt"

func numberOfCuts(n int) int {
	if n == 1 {
		return 0
	}
	if n&1 == 0 {
		return n / 2
	}
	return n
}

func main() {
	fmt.Println(numberOfCuts(1)) //0
	fmt.Println(numberOfCuts(6)) //3
	fmt.Println(numberOfCuts(5)) //5
}
