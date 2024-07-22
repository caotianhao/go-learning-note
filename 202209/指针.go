package main

import "fmt"

func main() {
	i := 10
	j := 20
	fmt.Println("the address of i is", &i)
	fmt.Println("the address of j is", &j)

	var p = &i
	fmt.Println("the address of i is", p, "i is", *p)

	*p = 11
	fmt.Println(i) //11
}
