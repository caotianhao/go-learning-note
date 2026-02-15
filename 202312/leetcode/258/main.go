package main

import "fmt"

func addDigits(num int) int {
	if num < 10 {
		return num
	}
	myNum, cnt, sum := num, 1, 0
	for myNum >= 10 {
		myNum /= 10
		cnt++
	}
	for i := 0; i < cnt; i++ {
		sum += num % 10
		num /= 10
	}
	if sum < 10 {
		return sum
	} else {
		return addDigits(sum)
	}
}

func main() {
	fmt.Println(addDigits(12356))
}
