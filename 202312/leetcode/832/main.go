package main

import "fmt"

func flipAndInvertImage(image [][]int) [][]int {
	l := len(image)
	for i := 0; i < l; i++ {
		image[i] = reverse832(image[i])
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			image[i][j] ^= 1
		}
	}
	return image
}

func reverse832(arr []int) []int {
	l := len(arr)
	for i := 0; i < l/2; i++ {
		arr[i], arr[l-1-i] = arr[l-1-i], arr[i]
	}
	return arr
}

func main() {
	arr := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
	fmt.Println(flipAndInvertImage(arr))
}
