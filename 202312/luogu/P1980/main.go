package main

import (
	"fmt"
	"strconv"
)

func main() {
	n, num := 0, 0
	_, _ = fmt.Scan(&n, &num)
	b := strconv.Itoa(num)
	count := 0
	for i := 1; i <= n; i++ {
		t := strconv.Itoa(i)
		for _, v := range t {
			if string(v) == b {
				count++
			}
		}
	}
	fmt.Println(count)
}
