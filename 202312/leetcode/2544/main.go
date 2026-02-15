package main

import "fmt"

func alternateDigitSum(n int) (sum int) {
	flag := 1
	for _, v := range reverse2544(divide2544(n)) {
		sum += flag * v
		flag = -flag
	}
	return
}

func divide2544(n int) (sli []int) {
	for n != 0 {
		sli = append(sli, n%10)
		n /= 10
	}
	return
}

func reverse2544(a []int) []int {
	l := len(a)
	for i := 0; i < l/2; i++ {
		a[i], a[l-1-i] = a[l-1-i], a[i]
	}
	return a
}

func main() {
	fmt.Println(alternateDigitSum(15678126))
}
