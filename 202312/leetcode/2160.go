package main

import "fmt"

func minimumSum(num int) int {
	numMap, numSlice := make(map[int]int), make([]int, 0)
	for num != 0 {
		numMap[num%10]++
		num /= 10
	}
	for i := 9; i >= 0; i-- {
		if numMap[i] != 0 {
			for j := 0; j < numMap[i]; j++ {
				numSlice = append(numSlice, i)
			}
		}
	}
	return numSlice[0] + numSlice[1] + 10*(numSlice[2]+numSlice[3])
}

func main() {
	fmt.Println(minimumSum(2239))
}
