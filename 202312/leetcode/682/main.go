package main

import (
	"fmt"
	"strconv"
)

func calPoints(operations []string) int {
	scoreSlice, score := make([]int, 0), 0
	for _, v := range operations {
		if v == "C" {
			scoreSlice = scoreSlice[:len(scoreSlice)-1]
		} else if v == "D" {
			t2 := scoreSlice[len(scoreSlice)-1]
			scoreSlice = append(scoreSlice, 2*t2)
		} else if v == "+" {
			t31 := scoreSlice[len(scoreSlice)-1]
			t32 := scoreSlice[len(scoreSlice)-2]
			scoreSlice = append(scoreSlice, t31+t32)
		} else {
			t4, _ := strconv.ParseInt(v, 10, 64)
			scoreSlice = append(scoreSlice, int(t4))
		}
	}
	for _, v := range scoreSlice {
		score += v
	}
	return score
}

func main() {
	fmt.Println(calPoints([]string{"5", "2", "C", "D", "+"}))
	fmt.Println(calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"}))
	fmt.Println(calPoints([]string{"1"}))
}
