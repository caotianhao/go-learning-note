package main

import "fmt"

func isHappy(n int) bool {
	mySlice, sum := make([]int, 0), 0
	for n != 0 {
		mySlice = append(mySlice, n%10)
		n /= 10
	}
	for i := 0; i < len(mySlice); i++ {
		sum += mySlice[i] * mySlice[i]
	}
	switch sum {
	case 1, 7, 10:
		return true
	case 2, 3, 4, 5, 6, 8, 9:
		return false
	}
	return isHappy(sum)
}

func main() {
	for i := 1; i < 123309; i++ {
		if isHappy(i) == true {
			fmt.Printf("%d ", i)
		}
		if i%53 == 0 {
			fmt.Println()
		}
	}
}
