package main

import "fmt"

func kk14(n int) int {
	s := 0
	for n != 0 {
		s += n % 10
		n /= 10
	}
	return s
}

func main() {
	k := 0
	for n := 1; n <= 1000; n++ {
		if kk14(n)%10 == 0 {
			k++
			//fmt.Printf("%d\n", n)
		}
	}
	fmt.Println(k)
}
