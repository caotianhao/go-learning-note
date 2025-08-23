package main

import "fmt"

func selfDividingNumbers(left int, right int) []int {
	mySlice := make([]int, 0)
	for i := left; i <= right; i++ {
		if isWant(i) {
			mySlice = append(mySlice, i)
		}
	}
	return mySlice
}

func isWant(n int) bool {
	mySlice, myN := make([]int, 0), n
	for myN != 0 {
		mySlice = append(mySlice, myN%10)
		myN /= 10
	}
	//flag:=true
	for _, v := range mySlice {
		if v == 0 || n%v != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(selfDividingNumbers(44, 77))
}
