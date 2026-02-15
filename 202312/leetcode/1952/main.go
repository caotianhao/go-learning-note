package main

import "fmt"

func isThree(n int) bool {
	cnt := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			cnt++
		}
	}
	return cnt == 3
}

func main() {
	fmt.Println(isThree(166))
}
