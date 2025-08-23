package main

import "fmt"

func countEven(num int) int {
	cnt := 0
	for i := 2; i <= num; i++ {
		if everySum2180(i)%2 == 0 {
			cnt++
		}
	}
	return cnt
}

func everySum2180(n int) int {
	sum := 0
	for n != 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func main() {
	fmt.Println(countEven(19))
}
