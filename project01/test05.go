package main

import (
	"fmt"
	"math/big"
)

var one = big.NewInt(1)

//var two = big.NewInt(2)

func L(x, n *big.Int) *big.Int {
	xMinusOne := new(big.Int).Sub(x, one)
	return new(big.Int).Div(xMinusOne, n)
}

func main() {

	x, y := new(big.Int), new(big.Int)

	// 随机产生8位素数p和q
	//p, _ := rand.Prime(rand.Reader, 6)
	//q, _ := rand.Prime(rand.Reader, 6)
	p := big.NewInt(11)
	q := big.NewInt(13)

	// n=p*q
	n := new(big.Int).Mul(p, q)
	// nn=n*n
	nn := new(big.Int).Mul(n, n)
	pMinusOne := new(big.Int).Sub(p, one)
	qMinusOne := new(big.Int).Sub(q, one)
	// phi(n)=(p-1)(q-1)
	phi := new(big.Int).Mul(pMinusOne, qMinusOne)
	// gcd=gcd(p-1,q-1)
	gcd := new(big.Int).GCD(x, y, pMinusOne, qMinusOne)
	// lambda=lcm(p-1,q-1)
	lambda := new(big.Int).Div(phi, gcd)

	//fmt.Println("p =", p)
	//fmt.Println("q =", q)
	//fmt.Println("n =", n)
	//fmt.Println("nn =", nn)
	//fmt.Println("phi =", phi)
	//fmt.Println("gcd =", gcd)
	//fmt.Println("lambda =", lambda)

	// g=n+1
	g := new(big.Int).Add(n, one)
	//fmt.Println("g =", g)

	//r, _ := rand.Int(rand.Reader, n)
	r := big.NewInt(7)
	//fmt.Println("r =", r)

	//m, _ := rand.Int(rand.Reader, new(big.Int).Div(n, two))
	m := big.NewInt(5)
	//fmt.Println("m =", m)

	gm := new(big.Int).Exp(g, m, nn)
	fmt.Println("g^m mod n^2 =", gm)

	rn := new(big.Int).Exp(r, n, nn)
	fmt.Println("r^n mod n^2 =", rn)

	c := new(big.Int).Mod(new(big.Int).Mul(gm, rn), nn)
	fmt.Println("c =", c)

	cLambda := new(big.Int).Exp(c, lambda, nn)
	fmt.Println("c^lambda mod n^2 =", cLambda)

	gLambda := new(big.Int).Exp(g, lambda, nn)
	fmt.Println("g^lambda mod n^2 =", gLambda)

	temp := new(big.Int).Div(L(cLambda, n), L(gLambda, n))
	fmt.Println("temp =", temp)

	tempN := new(big.Int).Mod(temp, n)
	fmt.Println("m =", tempN)

	//aaa := big.NewInt(111)
	//bbb := big.NewInt(7)
	//
	//fmt.Println(L(aaa, bbb))
}
