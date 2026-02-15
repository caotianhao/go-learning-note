package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

var zero3 = big.NewInt(0)
var one3 = big.NewInt(1)

// the whole key for paillier
type key struct {
	p      *big.Int
	q      *big.Int
	n      *big.Int
	g      *big.Int
	r      *big.Int
	nn     *big.Int
	p1     *big.Int
	q1     *big.Int
	gcd    *big.Int
	lambda *big.Int
	invLg  *big.Int
}

// function L(x)=(x-1)/n
func l3(x *big.Int, n *big.Int) *big.Int {
	return new(big.Int).Div(new(big.Int).Sub(x, one3), n)
}

func keyGen(bits int) *key {
	//x,y is just for GCD
	x := new(big.Int)
	y := new(big.Int)

	//choose p and q randomly
	p, _ := rand.Prime(rand.Reader, bits/2)
	q, _ := rand.Prime(rand.Reader, bits/2)

	//let p!=q
	flag := p.Cmp(q)
	for flag == 0 {
		p, _ = rand.Prime(rand.Reader, bits)
		q, _ = rand.Prime(rand.Reader, bits)
		flag = p.Cmp(q)
		if flag != 0 {
			break
		}
	}

	//n=p*q
	n := new(big.Int).Mul(p, q)

	//let g=n+1
	g := new(big.Int).Add(n, one3)

	//choose r randomly
	r, _ := rand.Int(rand.Reader, n)

	//let r!=0
	flag1 := r.Cmp(zero3)
	for flag1 == 0 {
		r, _ = rand.Int(rand.Reader, n)
		flag1 = r.Cmp(zero3)
		if flag != 0 {
			break
		}
	}

	//nn=n^2,p1=p-1,q1=q-1
	nn := new(big.Int).Mul(n, n)
	p1 := new(big.Int).Sub(p, one3)
	q1 := new(big.Int).Sub(q, one3)

	//find gcd(p-1,q-1) to get lcm(p-1,q-1)=lambda
	gcd := new(big.Int).GCD(x, y, p1, q1)
	lambda := new(big.Int).Div(new(big.Int).Mul(p1, q1), gcd)

	//g^lambda mod n^2, L it and L's inverse
	gLambdaModNn := new(big.Int).Exp(g, lambda, nn)
	lg := l3(gLambdaModNn, n)
	invLg := new(big.Int).ModInverse(lg, n)

	//return the whole key
	return &key{
		p:      p,
		q:      q,
		n:      n,
		g:      g,
		r:      r,
		nn:     nn,
		p1:     p1,
		q1:     q1,
		gcd:    gcd,
		lambda: lambda,
		invLg:  invLg,
	}
}

func enCry(message *big.Int, myKey key) *big.Int {
	//c=g^m*r^n mod n^2
	c1 := new(big.Int).Exp(myKey.g, message, myKey.nn)
	c2 := new(big.Int).Exp(myKey.r, myKey.n, myKey.nn)
	c := new(big.Int).Mod(new(big.Int).Mul(c1, c2), myKey.nn)
	return c
}

func deCry(cipher *big.Int, myKey key) *big.Int {
	//decry the cipher
	cLambdaModNn := new(big.Int).Exp(cipher, myKey.lambda, myKey.nn)
	lc := l3(cLambdaModNn, myKey.n)
	d := new(big.Int).Mod(new(big.Int).Mul(lc, myKey.invLg), myKey.n)
	return d
}

func main() {
	//test the time cost of keyGen which p and q are 1536 bits
	timeKeyGen0 := time.Now()
	k := keyGen(3072)
	timeKeyGen := time.Since(timeKeyGen0)
	fmt.Println("keyGen:", timeKeyGen)

	//make an array to save 1000 random messages
	var message [1000]*big.Int
	var messageCipher [1000]*big.Int

	//test encry and decry cost
	var enCryTime time.Duration = 0
	var deCryTime time.Duration = 0

	//forty := big.NewInt(40)

	//1000 random messages
	for i := 0; i < 1000; i++ {
		message[i], _ = rand.Int(rand.Reader, k.n)
		//message[i], _ = rand.Prime(rand.Reader, 1024)
		//message[i], _ = rand.Int(rand.Reader, forty)
	}

	//1000 times encry
	for i := 0; i < 1000; i++ {
		timeEnCry0 := time.Now()
		messageCipher[i] = enCry(message[i], *k)
		timeEnCry := time.Since(timeEnCry0)
		enCryTime += timeEnCry
	}

	//the average time of 1000 times encry
	fmt.Println("enCry 1000 cost_ave:", enCryTime/1000)

	//1000 times decry
	for i := 0; i < 1000; i++ {
		timeDeCry0 := time.Now()
		deCry(messageCipher[i], *k)
		timeDeCry := time.Since(timeDeCry0)
		deCryTime += timeDeCry
	}

	//the average time of 1000 times decry
	fmt.Println("deCry 1000 cost_ave:", deCryTime/1000)
}
