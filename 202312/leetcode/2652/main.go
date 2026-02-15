package main

import "fmt"

func sumOfMultiples(n int) (sum int) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 || i%7 == 0 {
			sum += i
		}
	}
	return
}

func main() {
	fmt.Println(sumOfMultiples(96))
}
