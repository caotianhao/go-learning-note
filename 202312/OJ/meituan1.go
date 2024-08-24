package main

import "fmt"

func main() {
	var (
		q    int
		m, x int64
	)
	_, _ = fmt.Scanf("%d", &q)
	for i := 0; i < q; i++ {
		_, _ = fmt.Scanf("%d %d", &m, &x)
		if x > m {
			if x%m == 0 {
				fmt.Printf("%d\n", m)
			} else {
				fmt.Printf("%d\n", (x%m+m)%m)
			}
		} else {
			fmt.Printf("%d\n", x)
		}
	}
}
