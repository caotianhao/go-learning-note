package main

import "fmt"

func winRed(a, b int) string {
	if (a+b-2)&1 == 1 {
		return "Yes"
	}
	return "No"
}

func main() {
	n := 0
	_, _ = fmt.Scanf("%d", &n)
	a, b := 0, 0
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d %d", &a, &b)
		fmt.Println(winRed(a, b))
	}
}
