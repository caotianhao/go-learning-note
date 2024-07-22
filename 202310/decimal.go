package main

import (
	"fmt"
	"math/big"
)

func main() {
	f1 := big.NewFloat(0.1)
	f2 := big.NewFloat(0.2)
	f := new(big.Float).Add(f1, f2)
	fmt.Println(f)

	r1 := new(big.Rat).SetFloat64(1.2)
	r2 := new(big.Rat).SetFloat64(3.4)
	r := new(big.Rat).Add(r1, r2)
	fmt.Println(r.String())

	c1, _ := new(big.Int).SetString("20716558285904281", 10)
	c2, _ := new(big.Int).SetString("4503599627370496", 10)
	c := new(big.Float).Quo(new(big.Float).SetInt(c1), new(big.Float).SetInt(c2))
	fmt.Println(c)
}
