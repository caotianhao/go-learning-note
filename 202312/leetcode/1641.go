package main

import "fmt"

func countVowelStrings(n int) int {
	preSum := make([][]int, n+1)
	preSum[0], preSum[1] = []int{}, []int{1, 1, 1, 1, 1}
	for i := 2; i <= n; i++ {
		for j := 0; j < 5; j++ {
			tmp := 0
			for k := 0; k <= j; k++ {
				tmp += preSum[i-1][k]
			}
			preSum[i] = append(preSum[i], tmp)
		}
	}
	return preSum1641(preSum[n])
}

func preSum1641(arr []int) (sum int) {
	for _, v := range arr {
		sum += v
	}
	return
}

func main() {
	fmt.Println(countVowelStrings(33)) //66045
	fmt.Println(countVowelStrings(1))  //66045
}
