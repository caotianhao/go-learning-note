package main

import (
	"fmt"
	"math"
)

func myCeil(n float64) int {
	return int(math.Ceil(n))
}

func main() {
	n, m := 0, 0
	k, c := 0, 0
	p := 0
	_, _ = fmt.Scanf("%d %d", &n, &m)
	people := make([]int, m)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d %d", &k, &c)
		money := myCeil(float64(c) / float64(k))
		for j := 0; j < k-1; j++ {
			_, _ = fmt.Scanf("%d", &p)
			people[p-1] += money
		}
	}
	for _, v := range people {
		fmt.Printf("%d ", v)
	}
}
