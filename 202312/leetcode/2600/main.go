package main

import "fmt"

func kItemsWithMaximumSum(numOnes, numZeros, numNegOnes, k int) int {
	if k <= numOnes {
		return k + numNegOnes - numNegOnes
	} else if k > numOnes && k <= numOnes+numZeros {
		return numOnes
	} else {
		return numOnes*2 - k + numZeros
	}
}

func main() {
	fmt.Println(kItemsWithMaximumSum(3, 2, 0, 2))
	fmt.Println(kItemsWithMaximumSum(3, 2, 6, 4))
	fmt.Println(kItemsWithMaximumSum(3, 2, 6, 1))
	fmt.Println(kItemsWithMaximumSum(3, 2, 6, 7))
}
