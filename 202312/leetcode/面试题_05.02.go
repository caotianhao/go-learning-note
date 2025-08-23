package main

import (
	"fmt"
	"math"
)

func printBin(num float64) string {
	ret, cnt := "0.", 3
	for num != 1 {
		num *= 2
		if num > 1 {
			ret += "1"
			num -= 1
			cnt++
		} else if num < 1 {
			ret += "0"
			cnt++
		}
		if cnt > 32 {
			return "ERROR"
		}
	}
	return ret + "1"
}

func main() {
	fmt.Println(printBin(0.625)) //0.101
	fmt.Println(printBin(0.5))
	fmt.Println(printBin(0.1))             //err
	fmt.Println(printBin(0.125))           //0.001
	fmt.Println(printBin(0.0625))          //0.0001
	fmt.Println(printBin(math.Pow(2, -4))) //0.0001
	fmt.Println(printBin(math.Pow(2, -5))) //0.00001  7位
	fmt.Println(printBin(math.Pow(2, -30)))
	fmt.Println(printBin(math.Pow(2, -31)))
	fmt.Println(printBin(math.Pow(2, -32)))
	fmt.Println(len(printBin(math.Pow(2, -30))))
}
