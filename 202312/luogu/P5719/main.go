package main

import "fmt"

func main() {
	var n, k, sa, ca, sb, cb int
	_, _ = fmt.Scanf("%d %d", &n, &k)
	for i := 1; i <= n; i++ {
		if i%k == 0 {
			sa += i
			ca++
		} else {
			sb += i
			cb++
		}
	}
	fmt.Printf("%.1f %.1f", float64(sa)/float64(ca), float64(sb)/float64(cb))
}
