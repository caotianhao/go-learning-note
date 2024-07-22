package main

import (
	"fmt"
	"math"
)

func main() {

	p := 11
	q := 13
	n := p * q
	//phi := (p-1)*(q-1)
	//lambda := 60
	nn := n * n
	g := n + 1
	r := 7
	m := 20
	c0 := math.Pow(float64(g), float64(m)) * math.Pow(float64(r), float64(n))
	c := int(c0) % nn
	fmt.Println(c)
}
