package main

import "fmt"

func getMaximumGenerated(n int) (max int) {
	if n < 2 {
		return n
	}
	arr := make([]int, n+1)
	arr[0], arr[1] = 0, 1
	for i := 2; i <= n; i++ {
		if i&1 == 0 {
			arr[i] = arr[i/2]
		} else {
			arr[i] = arr[i/2] + arr[i/2+1]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}
	return
}

func main() {
	fmt.Println(getMaximumGenerated(18))
}
