package main

import "fmt"

func main() {
	var h, x, s int64
	_, _ = fmt.Scanf("%d %d", &h, &x)
	snowBall := make([]int64, 0)
	for i := int64(0); i < h; i++ {
		_, _ = fmt.Scanf("%d", &s)
		snowBall = append(snowBall, s)
	}
	snowMod := make([]int64, 0)
	mod := int64(1e9 + 7)
	phi := x
	for i := 0; i < 100; i++ {
		snowMod = append(snowMod, phi%mod)
		phi = phi * x % mod
	}
	t := int64(0)
	for i := int64(0); i < h; i++ {
		t += snowBall[i] * snowMod[i] % mod
		t %= mod
	}
	fmt.Println(t)
}
