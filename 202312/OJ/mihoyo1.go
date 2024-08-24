package main

import "fmt"

func main() {
	n := 0
	_, _ = fmt.Scanf("%d", &n)
	a, b, aa, bb := 0, 0, make([]int, 0), make([]int, 0)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &a)
		aa = append(aa, a)
	}
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &b)
		bb = append(bb, b)
	}
	sa, sb := 0, 0
	for i := 0; i < n; i++ {
		sa += aa[i]
		sb += bb[i]
	}
	maxN := -1
	for i := 0; i < n; i++ {
		t := (sa - aa[i]) ^ sb
		if t > maxN {
			maxN = t
		}
	}
	for i := 0; i < n; i++ {
		t := (sb - bb[i]) ^ sa
		if t > maxN {
			maxN = t
		}
	}
	fmt.Printf("%d", maxN)
}
