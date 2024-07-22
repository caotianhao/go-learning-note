package main

import (
	"fmt"
	"math/big"
)

func main() {
	p := big.NewInt(35)
	q := big.NewInt(351)
	r := big.NewInt(3522)
	var c = 66
	var arr = [...]*big.Int{p, q, r}
	fmt.Println(c)
	fmt.Println(arr[2])
}
