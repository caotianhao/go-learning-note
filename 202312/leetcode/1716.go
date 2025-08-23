package main

import "fmt"

func totalMoney(n int) int {
	money := 0
	for i := 0; i < n; i++ {
		money += i%7 + 1 + i/7
		fmt.Println(i, money)
	}
	return money
}

func main() {
	fmt.Println(totalMoney(8))
}
