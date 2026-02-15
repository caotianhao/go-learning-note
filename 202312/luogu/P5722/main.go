package main

import "fmt"

func getSum5722(n int) int {
	if n == 1 {
		return 1
	}
	return n + getSum5722(n-1)
}

func main() {
	var a int
	_, _ = fmt.Scanf("%d", &a)
	fmt.Printf("%d", getSum5722(a))
}
