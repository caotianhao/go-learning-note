package main

import (
	"fmt"
	"math"
)

var squareNumberBilibili2 []int

func init() {
	var i int
	for ; i < 33333; i++ {
		squareNumberBilibili2 = append(squareNumberBilibili2, i*i)
	}
}

func toSquareNumber(x int) int {
	for _, v := range squareNumberBilibili2 {
		if x == v {
			return 0
		}
	}
	sqrt := math.Sqrt(float64(x))
	left := int(math.Floor(sqrt))
	right := int(math.Ceil(sqrt))
	o := x % 2
	if right%2 == o && left%2 == o {
		return min((right*right-x)/2, (x-left*left)/2)
	} else if right%2 != o && left%2 == o {
		return min(((right+1)*(right+1)-x)/2, (x-left*left)/2)
	} else if right%2 == o && left%2 != o {
		return min((right*right-x)/2, (x-(left-1)*(left-1))/2)
	} else {
		return min(((right+1)*(right+1)-x)/2, (x-(left-1)*(left-1))/2)
	}
}

func main() {
	var x int
	_, _ = fmt.Scanf("%d", &x)
	fmt.Printf("%d", toSquareNumber(x))
}
