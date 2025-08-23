package main

import "fmt"

func xorOperation(n int, start int) int {
	mySlice := make([]int, 0)
	for i := 0; i < n; i++ {
		mySlice = append(mySlice, start+2*i)
	}
	xor := mySlice[0]
	for i := 1; i < n; i++ {
		xor ^= mySlice[i]
	}
	return xor
}

func main() {
	fmt.Println(xorOperation(255, 12))
}
