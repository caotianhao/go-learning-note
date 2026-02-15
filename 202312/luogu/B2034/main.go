package main

import "fmt"

func get2Mi(n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return 2
	} else {
		return 2 * get2Mi(n-1)
	}
}

func main() {
	var a int
	_, _ = fmt.Scanf("%d", &a)
	fmt.Printf("%d", get2Mi(a))
}
