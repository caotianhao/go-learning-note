package main

import "fmt"

func arrangeCoins(n int) int {
	//sum := make([]int, 0)
	//sum = append(sum, 0)
	//for i := 1; i <= 65536; i++ {
	//	sum = append(sum, i*(i+1)/2)
	//}
	for i := 0; i < 65537; i++ {
		if n < i*(i+1)/2 {
			return i - 1
		}
	}
	return -1
}

func main() {
	fmt.Println(arrangeCoins(10))
}
