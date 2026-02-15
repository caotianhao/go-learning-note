package main

import "fmt"

func main() {
	var num float64
	_, _ = fmt.Scanf("%f", &num)
	tenPlus := int(num * 10)
	one, two, three, four := tenPlus%10, tenPlus/10%10, tenPlus/10/10%10, tenPlus/10/10/10%10
	fmt.Printf("%.3f", float64(one)+float64(two)*0.1+float64(three)*0.01+float64(four)*0.001)
}
