package main

import "fmt"

func subtractProductAndSum(n int) int {
	mySlice, myNum, pro, sum := make([]int, 0), n, 1, 0
	for myNum != 0 {
		mySlice = append(mySlice, myNum%10)
		myNum /= 10
	}
	for _, v := range mySlice {
		pro *= v
		sum += v
	}
	return pro - sum
}

func main() {
	fmt.Println(subtractProductAndSum(16))
	fmt.Println(subtractProductAndSum(1611))
}
