package main

import "fmt"

const (
	gold2red    = 28080
	red2normal  = 58760
	normal2best = 70200
	base        = 2300
)

const (
	fenjie = 258720
	have   = 270000 + 70*base
	now    = fenjie + have
)

func main() {
	normal3 := 3 * (gold2red + red2normal)
	goods := 2 * (gold2red + red2normal + normal2best)
	weapon := red2normal + normal2best

	all := normal3 + goods + weapon

	fmt.Println((all - now) / base)
}

// prepare 22
// but after calculation
// change 23
