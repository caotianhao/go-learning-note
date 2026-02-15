package main

import "fmt"

func main() {
	var n, a, b, mod int64
	_, _ = fmt.Scanf("%d %d %d %d", &n, &a, &b, &mod)
	edge := make([]int64, n-1)
	var i int64
	for ; i < n-1; i++ {
		edge[i] = i + int64(2) // 2 3 4 5 6 7
	}
	var tr int64
	trSlice := make([]int64, 0)
	trMap := make(map[int64]struct{})
	tree := make(map[int64][]int64)
	var j int64
	for ; j < n-1; j++ {
		_, _ = fmt.Scanf("%d", &tr)
		trSlice = append(trSlice, tr) // 1 1 2 2 3 6
		trMap[tr] = struct{}{}
		tree[tr] = append(tree[tr], edge[j])
	}
	h := make([]int64, n+1)
	h[0] = -1
	k := n
	for ; k >= 1; k-- {
		if _, ok := trMap[k]; ok {
			aa := plus(h[tree[k][0]]+myPow(a, k, mod), b, mod)
			if len(tree[k]) == 1 {
				// h6 = (h7+a^6)*b
				h[k] = aa
			} else if len(tree[k]) == 2 {
				bb := plus(h[tree[k][1]]+myPow(a, k, mod), b, mod)
				h[k] = (aa + bb) % mod
			}
		}
	}
	fmt.Printf("%d", h[1])
}

func myPow(x, y, mod int64) int64 {
	var r int64 = 1
	if y == 1 {
		return x % mod
	}
	for y != 0 {
		r *= x
		y--
		if r > mod {
			r = r % mod
		}
	}
	return r % mod
}

func plus(a, b, mod int64) int64 {
	var r int64
	for b != 0 {
		r += a
		if r > mod {
			r = r % mod
		}
		b--
	}
	return r
}
