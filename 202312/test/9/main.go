package main

import (
	"fmt"
	"math/big"
)

var (
	one = big.NewInt(1)
)

func factories(n *big.Int) *big.Int {
	res := one
	for i := one; i.Cmp(n) <= 0; i = new(big.Int).Add(i, one) {
		res = new(big.Int).Mul(res, i)
	}
	return res
}

func main() {
	n := big.NewInt(100)
	fmt.Println(factories(n))
}
