package main

import (
	"fmt"
	"math/big"
)

func main() {
	p := big.NewInt(469232994506)
	q := big.NewInt(366024)
	//r := big.NewInt(253)
	zero := big.NewInt(0)
	a := new(big.Int).Exp(p, q, zero)
	fmt.Println(a)
}
