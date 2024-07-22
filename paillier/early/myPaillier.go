package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

// encry 60-70 ms
// decry 10-  ms

var zero = big.NewInt(0)
var one = big.NewInt(1)

func l(x *big.Int, n *big.Int) *big.Int {
	return new(big.Int).Div(new(big.Int).Sub(x, one), n)
}

func main() {
	// x,y for gcd
	time0 := time.Now()
	x := new(big.Int)
	y := new(big.Int)
	// the bit of the sk
	bits := 3072
	// choose p and q randomly
	p, _ := rand.Prime(rand.Reader, bits)
	q, _ := rand.Prime(rand.Reader, bits)
	// if p==q , choose again
	flag := p.Cmp(q)
	for flag == 0 {
		p, _ = rand.Prime(rand.Reader, bits)
		q, _ = rand.Prime(rand.Reader, bits)
		flag = p.Cmp(q)
		if flag != 0 {
			break
		}
	}

	//p:=big.NewInt(13)
	//q:=big.NewInt(11)
	//fmt.Println("p =", p, "q =", q)
	// n=pq
	n := new(big.Int).Mul(p, q)

	//time0 := time.Now()
	// g=n+1
	g := new(big.Int).Add(n, one)

	// random r!=0
	r, _ := rand.Int(rand.Reader, n)
	flag1 := r.Cmp(zero)
	for flag1 == 0 {
		r, _ = rand.Int(rand.Reader, n)
		flag1 = r.Cmp(zero)
		if flag != 0 {
			break
		}
	}

	//r := big.NewInt(31)
	//fmt.Println("r =", r)
	// nn=n*n
	nn := new(big.Int).Mul(n, n)
	//fmt.Println("nn =", nn)
	// p1=p-1,q1=q-1
	p1 := new(big.Int).Sub(p, one)
	q1 := new(big.Int).Sub(q, one)
	// gcd(p-1,q-1)
	gcd := new(big.Int).GCD(x, y, p1, q1)
	// lcm(p-1,q-1)
	lambda := new(big.Int).Div(new(big.Int).Mul(p1, q1), gcd)
	//fmt.Println("lambda =", lambda)
	// random message<n
	m, _ := rand.Int(rand.Reader, n)
	fmt.Println("m =", m)
	// time begin
	//time0 := time.Now()
	//m:=big.NewInt(59)

	// exactly choose g
	//g, _ := rand.Int(rand.Reader, nn)
	//flag2 := one.Cmp(l(new(big.Int).Mod(new(big.Int).Exp(g, lambda, zero), nn), n))
	//for ; flag2 != 0; {
	//	g, _ = rand.Int(rand.Reader, nn)
	//	flag2 = one.Cmp(l(new(big.Int).Mod(new(big.Int).Exp(g, lambda, zero), nn), n))
	//	if flag2 == 0 {
	//		break
	//	}
	//}

	// c=g^m*r^n mod n^2
	c1 := new(big.Int).Exp(g, m, nn)
	c2 := new(big.Int).Exp(r, n, nn)
	c := new(big.Int).Mod(new(big.Int).Mul(c1, c2), nn)
	//fmt.Println("c =", c)

	// c^lambda
	//cLambda := new(big.Int).Exp(c, lambda, zero)
	//fmt.Println("cLambda =", cLambda)
	// c^lambda mod n^2
	cLambdaModNn := new(big.Int).Exp(c, lambda, nn)
	//fmt.Println("cLambdaModNn =", cLambdaModNn)
	// L(c^lambda mod n^2)
	lc := l(cLambdaModNn, n)
	//fmt.Println("lc =", lc)

	//time0 := time.Now()
	// g^lambda
	//gLambda := new(big.Int).Exp(g, lambda, zero)
	// g^lambda mod n^2
	gLambdaModNn := new(big.Int).Exp(g, lambda, nn)
	// L(g^lambda mod n^2)
	lg := l(gLambdaModNn, n)
	//fmt.Println("lg =", lg)

	// (L(g^lambda mod n^2))^(-1)
	//invLg := new(big.Int).ModInverse(lg, n)
	//time00 := time.Now()
	xx := new(big.Int)
	yy := new(big.Int)
	gcdD := new(big.Int).GCD(xx, yy, n, lg)
	if zero.Cmp(yy) == 1 {
		yy = new(big.Int).Add(yy, n)
	}
	invLg := yy
	//time10 := time.Since(time00)
	//fmt.Println(time10)

	//fmt.Println("invLg =", invLg)
	// decry
	//time0 := time.Now()
	d := new(big.Int).Mod(new(big.Int).Mul(lc, invLg), n)
	// time end
	time1 := time.Since(time0)
	fmt.Println(bits, "bits time =", time1)
	fmt.Println("d =", d, gcdD)
	//fmt.Println(gcdD)
	//fmt.Println("p =", p, ",q =", q, ",n =", n, ",lambda =", lambda, ",g =", g, ",nn =", nn)
	//fmt.Println("c =", c, "lc =", lc, ",lg =", lg)

}
