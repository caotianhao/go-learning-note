package main

import "fmt"

func numberOfSteps(num int) int {
	cnt := 0
	for num != 0 {
		if num%2 == 0 {
			num /= 2
			cnt++
		} else {
			num -= 1
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(numberOfSteps(26))
}
