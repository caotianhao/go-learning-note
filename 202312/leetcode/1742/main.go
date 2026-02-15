package main

import "fmt"

func countBalls(lowLimit int, highLimit int) int {
	map1, maxN := make(map[int]int), 0
	for i := lowLimit; i <= highLimit; i++ {
		map1[everySum(i)]++
	}
	for i, v := range map1 {
		if map1[i] > maxN {
			maxN = v
		}
	}
	return maxN
}

func everySum(n int) int {
	sum := 0
	for n != 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func main() {
	fmt.Println(countBalls(12, 36))
}
