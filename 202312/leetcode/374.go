package main

import "fmt"

func guessNumber(n int) int {
	l, r, mid := 1, n, -1
	for l <= r {
		mid = (l + r) / 2
		if guess(mid) == 1 {
			l = mid + 1
		} else if guess(mid) == -1 {
			r = mid - 1
		} else {
			return mid
		}
	}
	return mid
}

func guess(a int) int {
	return a - 1
}

func main() {
	fmt.Println(guessNumber(1222))
}
