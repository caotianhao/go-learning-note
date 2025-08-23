package main

import (
	"fmt"
	"math"
)

func reverseBits(num uint32) uint32 {
	var ret uint32
	for i := 0; i < 32; i++ {
		if 1<<i&num != 0 {
			ret += uint32(math.Pow(2.0, float64(31-i)))
		}
	}
	return ret
}

func main() {
	fmt.Println(reverseBits(4294967293)) //3221225471
	fmt.Println(reverseBits(43261596))   //964176192
}
