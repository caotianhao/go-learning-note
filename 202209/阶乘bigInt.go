package main

import (
	"fmt"
	"math/big"
)

var myFactorialMapBigInt = make(map[*big.Int]*big.Int)

func factorial2(n *big.Int) {
	var res = big.NewInt(1)
	var one = big.NewInt(1)
	for i := big.NewInt(1); i.Cmp(n) <= 0; i = new(big.Int).Add(i, one) {
		res = new(big.Int).Mul(res, i)
	}
	myFactorialMapBigInt[n] = res
}

func main() {
	var one = big.NewInt(1)
	//输入 k，求 1~k 的阶乘
	var k int64
	_, _ = fmt.Scanln(&k)
	for i := big.NewInt(1); i.Cmp(big.NewInt(k)) <= 0; i = new(big.Int).Add(i, one) {
		factorial2(i)
	}
	for i, v := range myFactorialMapBigInt {
		fmt.Printf("%d!\t= %d\n", i, v)
	}
}
