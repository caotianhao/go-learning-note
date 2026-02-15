package main

import "fmt"

func accountBalanceAfterPurchase(purchaseAmount int) int {
	m := 11
	q := 0
	for i := 0; i <= 100; i += 10 {
		t := abs2806(i - purchaseAmount)
		if t <= m {
			m = t
			q = i
		}
	}
	return 100 - q
}

func abs2806(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func main() {
	fmt.Println(accountBalanceAfterPurchase(9))
	fmt.Println(accountBalanceAfterPurchase(2))
}
