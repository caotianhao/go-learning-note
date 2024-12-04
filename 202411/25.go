package main

import (
	"fmt"
)

const (
	XIAO_HONG_MAO = 2693.3
	END_OF_LIFE   = 2789.38

	GOOD_FOUR_YEARS  = 8954.5
	END_OF_THE_WORLD = 68377.84

	BEGIN_OF_THE_WORLD = 7.19
)

func main() {
	don := XIAO_HONG_MAO + END_OF_LIFE + GOOD_FOUR_YEARS + END_OF_THE_WORLD + BEGIN_OF_THE_WORLD
	fmt.Println(don)
	fmt.Println(int(don / 0.005140))
}
