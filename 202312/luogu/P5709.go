package main

import (
	"fmt"
	"math"
)

func main() {
	var m, t, s int
	_, _ = fmt.Scanf("%d %d %d", &m, &t, &s)
	if t == 0 {
		fmt.Printf("%d", 0)
	} else {
		tmp := m - int(math.Ceil(float64(s)/float64(t)))
		if tmp >= 0 {
			fmt.Printf("%d", tmp)
		} else {
			fmt.Printf("%d", 0)
		}
	}
}
