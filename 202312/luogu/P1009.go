package main

import (
	"fmt"
	"math/big"
)

var one = big.NewInt(1)
var zero = big.NewInt(0)

func jc1009(n *big.Int) *big.Int {
	res := one
	for i := n; i.Cmp(one) == 1; i = new(big.Int).Sub(i, one) {
		res = new(big.Int).Mul(res, i)
	}
	return res
}

func main() {
	var n int64
	sum := zero
	_, _ = fmt.Scanf("%d", &n)
	num := big.NewInt(n)
	for i := num; i.Cmp(one) >= 0; i = new(big.Int).Sub(i, one) {
		sum = new(big.Int).Add(sum, jc1009(i))
	}
	fmt.Printf("%d", sum)
}
