package main

import "fmt"

func decode(encoded []int, first int) []int {
	mySlice := make([]int, 0)
	temp := first ^ encoded[0]
	mySlice = append(mySlice, first, temp)
	for i := 1; i < len(encoded); i++ {
		mySlice = append(mySlice, temp^encoded[i])
		temp = temp ^ encoded[i]
	}
	return mySlice
}

func main() {
	arr := []int{1, 2, 3}
	first := 1
	fmt.Println(decode(arr, first))
}
