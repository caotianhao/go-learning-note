package main

import "fmt"

func main() {
	for i := 1; ; i++ {
		tmp := i * 4000
		if tmp%4500 == 0 && tmp%5000 == 0 && tmp%5500 == 0 && tmp%6500 == 0 && tmp%7260 == 0 {
			fmt.Println(i, tmp)
			break
		}
	}
}
