package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scanf("%d", &n)
	cnt := 1
	for i := n; i >= 1; i-- {
		for j := 0; j < i; j++ {
			fmt.Printf("%02d", cnt)
			cnt++
		}
		fmt.Printf("\n")
	}
}
