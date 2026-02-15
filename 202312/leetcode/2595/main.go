package main

import "fmt"

func evenOddBit(n int) []int {
	even, odd := 0, 0
	for i := 0; i < 10; i++ {
		if (n>>i)&1 == 1 {
			if i%2 == 0 {
				even++
			} else {
				odd++
			}
		}
	}
	return []int{even, odd}
}

func main() {
	fmt.Println(evenOddBit(17))
}
