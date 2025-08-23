package main

import "fmt"

func hammingDistance(x int, y int) int {
	xx, yy, cnt := decToBin461(x), decToBin461(y), 0
	for i := 0; i < 31; i++ {
		if xx[i] != yy[i] {
			cnt++
		}
	}
	return cnt
}

func decToBin461(x int) []int {
	if x == 0 {
		slice4611 := make([]int, 31)
		return slice4611
	}
	slice461 := make([]int, 0)
	for x != 0 {
		slice461 = append(slice461, x%2)
		x /= 2
	}
	return addZero461(reverse461(slice461))
}

func reverse461(arr []int) []int {
	l := len(arr)
	for i := 0; i < l/2; i++ {
		arr[i], arr[l-i-1] = arr[l-i-1], arr[i]
	}
	return arr
}

func addZero461(arr []int) []int {
	l := len(arr)
	if l == 31 {
		return arr
	}
	slice461 := make([]int, 31-l)
	for _, v := range arr {
		slice461 = append(slice461, v)
	}
	return slice461
}

func main() {
	fmt.Println(hammingDistance(0, 1))
	fmt.Println(decToBin461(0))
}
