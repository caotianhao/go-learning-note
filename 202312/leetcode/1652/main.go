package main

import "fmt"

func decrypt(code []int, k int) []int {
	l := len(code)
	message := make([]int, l)
	if k == 0 {
		return message
	}
	for i := 0; i < l; i++ {
		message[i] = code[i]
	}
	if k > 0 {
		for i := 0; i < l; i++ {
			temp, cnt := 0, 0
			for j := i; cnt < k; j++ {
				temp += message[(j+1)%l]
				cnt++
			}
			code[i] = temp
		}
	}
	if k < 0 {
		for i := 0; i < l; i++ {
			temp, cnt := 0, 0
			for j := i; cnt < -k; j-- {
				temp += message[(l+j-1)%l]
				cnt++
			}
			code[i] = temp
		}
	}
	return code
}

func main() {
	fmt.Println(decrypt([]int{5, 7, 1, 4}, 3))
	fmt.Println(decrypt([]int{1, 2, 3, 4}, 0))
	fmt.Println(decrypt([]int{2, 4, 9, 3}, -2))
	fmt.Println(decrypt([]int{2}, -1))
}
