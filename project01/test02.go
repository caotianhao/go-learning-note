package main

import (
	"fmt"
	"math/big"
)

func main() {

	//im := big.NewInt(math.MaxInt64)
	//in := im
	//io := big.NewInt(1956)
	//ip := big.NewInt(1)
	//ip.Mul(im, in).Add(ip, im)
	//
	//fmt.Printf("Big Int: %v\n", im)
	//fmt.Printf("Big Int: %v\n", in)
	//fmt.Printf("Big Int: %v\n", io)
	//fmt.Printf("Big Int: %v\n", ip)

	//m := big.NewInt(math.MaxInt64)
	//n := m
	//o := big.NewInt(1956)
	//p := big.NewInt(1)
	//p.Mul(m, n).Add(p, o)
	//fmt.Println(m)
	//fmt.Printf("%v\n", p)

	a := big.NewInt(12345678910390)
	b := big.NewInt(999)
	c := big.NewInt(200)
	d := big.NewInt(4066666)
	a.Mul(a, b).Mul(a, c).Mul(a, d)
	fmt.Println(a)

}
