package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	// Here are some calculations with bigInts:
	im := big.NewInt(math.MaxInt64) // 9223372036854775807
	in := im
	io := big.NewInt(4349212256146964)
	ip := big.NewInt(1)
	//fmt.Println()
	//ip.Mul(im, in).Add(ip, im).Div(ip, io)
	//fmt.Printf("Big Int: %v\n", ip)
	fmt.Printf("im: %v\n", im)
	fmt.Printf("in: %v\n", in)
	fmt.Printf("io: %v\n", io)
	fmt.Printf("ip: %v\n", ip)
	ip.Mul(im, in).Add(im, ip).Div(ip, io).Div(ip, io)
	fmt.Printf("newip: %v\n", ip)

	// Here are some calculations with bigInts:
	//rm := big.NewRat(math.MaxInt64, 1956)
	//rn := big.NewRat(-1956, math.MaxInt64)
	//ro := big.NewRat(19, 56)
	//rp := big.NewRat(1111, 2222)
	//rq := big.NewRat(1, 1)
	//rq.Mul(rm, rn).Add(rq, ro).Mul(rq, rp)
	//fmt.Printf("Big Rat: %v\n", rq)
}

/* Output:
Big Int: 43492122561469640008497075573153004
Big Rat: -37/112
*/
