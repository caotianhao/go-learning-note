package main

import (
	"fmt"
	"math/big"
)

func main() {
	zero := big.NewInt(0)
	x := new(big.Int)
	y := new(big.Int)
	p := big.NewInt(24911235622227199)
	q := big.NewInt(7116636636633341)
	gcd := new(big.Int).GCD(x, y, p, q)
	if zero.Cmp(y) == 1 {
		y = new(big.Int).Add(y, p)
	}
	fmt.Println(gcd, "gcd:", y)
	modInv := new(big.Int).ModInverse(q, p)
	fmt.Println(modInv)
}
