package main

import "fmt"

func main() {
	var x, a, b, c, d int
	_, _ = fmt.Scanf("%d", &x)
	flag1, flag2 := x%2 == 0, x > 4 && x <= 12
	//fmt.Println(flag1, flag2)
	if flag1 && flag2 {
		a = 1
	} else {
		a = 0
	}
	if !(!flag1 && !flag2) {
		b = 1
	} else {
		b = 0
	}
	if (flag1 && !flag2) || (flag2 && !flag1) {
		c = 1
	} else {
		c = 0
	}
	d = b ^ 1
	fmt.Printf("%d %d %d %d", a, b, c, d)
}
