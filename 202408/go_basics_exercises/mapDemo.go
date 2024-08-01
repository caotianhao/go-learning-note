package main

import "fmt"

func f(a int) int {
	return a * a
}

func myMap(f func(i int) int, arr []int) (ret []int) {
	for _, v := range arr {
		ret = append(ret, f(v))
	}
	return
}

func main() {
	arr := make([]int, 8)
	for i := 0; i < 8; i++ {
		arr[i] = i + i*i
	}
	fmt.Println(myMap(f, arr))
}
