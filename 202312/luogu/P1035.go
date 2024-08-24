package main

import "fmt"

func main() {
	k := 0
	_, _ = fmt.Scan(&k)
	i := 1
	var now float64 = 0
	for {
		if now > float64(k) {
			break
		}
		now += float64(1) / float64(i)
		//fmt.Println(now, i)
		i++
	}
	fmt.Println(i - 1)
}
