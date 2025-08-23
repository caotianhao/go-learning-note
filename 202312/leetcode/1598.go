package main

import (
	"fmt"
	"math"
)

func minOperations1598(logs []string) int {
	opt := 0
	for _, v := range logs {
		if v == "../" {
			if opt > 0 {
				opt--
			} else {
				opt += 0
			}
		} else if v == "./" {
			opt += 0
		} else {
			opt++
		}
	}
	return int(math.Max(float64(opt), float64(0)))
}

func main() {
	fmt.Println(minOperations1598([]string{"d1/", "d2/", "../", "d21/", "./"}))
	fmt.Println(minOperations1598([]string{"d1/", "d2/", "./", "d3/", "../", "d31/"}))
	fmt.Println(minOperations1598([]string{"d1/", "../", "../", "../"}))
	fmt.Println(minOperations1598([]string{"./", "wz4/", "../", "mj2/", "../", "../", "ik0/", "il7/"}))
}
