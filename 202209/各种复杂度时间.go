package main

import (
	"fmt"
	"time"
)

func oN(n int) {
	a := 0
	for i := 0; i < n; i++ {
		a++
	}
}

func oN2(n int) {
	a := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a++
		}
	}
}

func oN3(n int) {
	a := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				a++
			}
		}
	}
}

func oN4(n int) {
	a := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				for l := 0; l < n; l++ {
					a++
				}
			}
		}
	}
}

func oN5(n int) {
	a := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				for l := 0; l < n; l++ {
					for m := 0; m < n; m++ {
						a++
					}
				}
			}
		}
	}
}

func oNLogN(n int) {
	a := 0
	for i := 0; i < n; i++ {
		for j := 1; j < n; j *= 2 {
			a++
		}
	}
}

func oLogN(n int) {
	a := 0
	for i := 1; i < n; i *= 2 {
		a++
	}
}

func main() {
	var i int
	_, _ = fmt.Scanln(&i)
	time0 := time.Now()
	//oN(i)
	//oN(i)
	//oN2(i)
	oN3(i)
	//oN4(i)
	//oN5(i)
	//oNLogN(i)
	//oLogN(i)
	time1 := time.Since(time0)
	fmt.Println(time1)
}
