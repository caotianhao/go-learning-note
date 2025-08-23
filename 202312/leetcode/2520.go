package main

import "fmt"

func countDigits(num int) (cnt int) {
	for _, v := range divide2520(num) {
		if num%v == 0 {
			cnt++
		}
	}
	return
}

func divide2520(n int) (sli2520 []int) {
	for n != 0 {
		sli2520 = append(sli2520, n%10)
		n /= 10
	}
	return
}

func main() {
	fmt.Println(countDigits(1248)) //4
}
