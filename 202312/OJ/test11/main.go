package main

import "fmt"

func main() {
	a := make(map[int]struct{})
	for i := 0; i <= 300000; i++ {
		a[i*i] = struct{}{}
	}
	for i := 0; i <= 300000; i++ {
		_, ok := a[i+100]
		_, ok2 := a[i+168]
		if ok && ok2 {
			fmt.Println(i)
		}
	}
}
