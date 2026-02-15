package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	var isTime byte
	_, _ = fmt.Scanf("%d %c", &a, &isTime)
	if isTime == 'y' {
		if a <= 1000 {
			fmt.Printf("13")
		} else {
			fmt.Printf("%d", 13+int(math.Ceil(float64(a-1000)/500))*4)
		}
	} else {
		if a <= 1000 {
			fmt.Printf("8")
		} else {
			fmt.Printf("%d", 8+int(math.Ceil(float64(a-1000)/500))*4)
		}
	}
}
