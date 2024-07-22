package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	zero := big.NewInt(0)
	p, _ := rand.Prime(rand.Reader, 2560)
	q, _ := rand.Prime(rand.Reader, 2559)
	//bit := big.NewInt(qq)
	//p, _ := rand.Int(rand.Reader, bit)
	//q, _ := rand.Int(rand.Reader, bit)
	x := new(big.Int)
	y := new(big.Int)
	time0 := time.Now()
	gcd := new(big.Int).GCD(x, y, p, q)
	if zero.Cmp(y) == 1 {
		y = new(big.Int).Add(y, p)
	}
	timeGcd := time.Since(time0)
	fmt.Println(gcd, y)
	fmt.Println(timeGcd)
}
