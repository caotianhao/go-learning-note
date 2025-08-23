package main

import "fmt"

func removeTrailingZeros(num string) string {
	for num[len(num)-1] == '0' {
		num = num[:len(num)-1]
	}
	return num
}

func main() {
	fmt.Println(removeTrailingZeros("654100413120"))
	fmt.Println(removeTrailingZeros("64000030"))
	fmt.Println(removeTrailingZeros("6400003"))
	fmt.Println(removeTrailingZeros("20"))
}
