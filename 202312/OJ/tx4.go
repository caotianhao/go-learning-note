package main

import "fmt"

func deLivePush(a, b int) int {
	return a - b
}

func main() {
	n := 0
	_, _ = fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		a, b := 0, 0
		_, _ = fmt.Scanf("%d %d", &a, &b)
		fmt.Printf("%d\n", deLivePush(a, b))
	}
}
