package main

import "fmt"

func main() {
	var n, k, res int
	_, _ = fmt.Scanf("%d %d", &n, &k)
	for n >= k {
		n -= k
		res += k
		if n >= 0 {
			n++
		}
	}
	res += n
	fmt.Println(res)
}
