package main

import "fmt"

func main() {
	n, k := 0, 0
	_, _ = fmt.Scanf("%d %d", &n, &k)
	i := 0
	ii := make([]int, 0)
	for j := 0; j < n; j++ {
		_, _ = fmt.Scanf("%d", &i)
		ii = append(ii, i)
	}
	if n == 2 {
		fmt.Printf("%d", ii[1])
	}
}
