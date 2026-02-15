package main

import "fmt"

func fraction(cont []int) []int {
	res := make([]int, 2)
	cont = reverseLCP02(cont)
	res[0], res[1] = 1, cont[0]
	for i := 1; i < len(cont); i++ {
		res[0] += cont[i] * res[1]
		res[0], res[1] = res[1], res[0]
	}
	res[0], res[1] = res[1], res[0]
	if res[0] < 0 && res[1] < 0 {
		res[0], res[1] = -res[0], -res[1]
	}
	gc := gcdLCP02(res[0], res[1])
	if gcdLCP02(res[0], res[1]) != 1 {
		res[0], res[1] = res[0]/gc, res[1]/gc
	}
	return res
}

func reverseLCP02(arr []int) []int {
	l := len(arr)
	for i := 0; i < l/2; i++ {
		arr[i], arr[l-i-1] = arr[l-i-1], arr[i]
	}
	return arr
}

func gcdLCP02(a, b int) int {
	for b != 0 {
		return gcdLCP02(b, a%b)
	}
	return a
}

func main() {
	fmt.Println(fraction([]int{3, 2, 0, 2}))
	fmt.Println(fraction([]int{1, 2, 3, 4}))
	fmt.Println(fraction([]int{0, 0, 3}))
	fmt.Println(fraction([]int{2, 2}))
}
