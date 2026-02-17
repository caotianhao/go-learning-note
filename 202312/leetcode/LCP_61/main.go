package main

import (
	"fmt"
)

func temperatureTrend(temperatureA []int, temperatureB []int) int {
	l, res := len(temperatureA), 0
	resA, resB := make([]int, l-1), make([]int, l-1)
	for i := 0; i < l-1; i++ {
		if temperatureA[i] < temperatureA[i+1] {
			resA[i] = 1
		} else if temperatureA[i] == temperatureA[i+1] {
			resA[i] = 0
		} else {
			resA[i] = -1
		}
	}
	for i := 0; i < l-1; i++ {
		if temperatureB[i] < temperatureB[i+1] {
			resB[i] = 1
		} else if temperatureB[i] == temperatureB[i+1] {
			resB[i] = 0
		} else {
			resB[i] = -1
		}
	}
	for i := 0; i < l-1; i++ {
		for j := i; j < l-1; j++ {
			if isTwoArraysEqual(resA[i:j+1], resB[i:j+1]) {
				res = max(res, j-i+1)
			} else {
				break
			}
		}
	}
	return res
}

func isTwoArraysEqual(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	//[1,-15,3,14,-1,4,35,36]
	//[-15,32,20,9,33,4,-1,-5]
	//[-14,7,-19,9,13,40,19,15,-18]
	//[3,16,28,32,25,12,13,-6,4]
	fmt.Println(temperatureTrend([]int{-14, 7, -19, 9, 13, 40, 19, 15, -18},
		[]int{3, 16, 28, 32, 25, 12, 13, -6, 4})) //1
	fmt.Println(temperatureTrend([]int{21, 18, 18, 18, 31},
		[]int{34, 32, 16, 16, 17})) //2
}
