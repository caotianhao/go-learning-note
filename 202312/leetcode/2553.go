package main

import "fmt"

func separateDigits(nums []int) (sli2553 []int) {
	for _, v := range nums {
		for _, v1 := range reverse2553(divide2553(v)) {
			sli2553 = append(sli2553, v1)
		}
	}
	return
}

func divide2553(n int) (sli2553 []int) {
	for n != 0 {
		sli2553 = append(sli2553, n%10)
		n /= 10
	}
	return
}

func reverse2553(arr []int) []int {
	l := len(arr)
	for i := 0; i < l/2; i++ {
		arr[i], arr[l-1-i] = arr[l-1-i], arr[i]
	}
	return arr
}

func main() {
	fmt.Println(separateDigits([]int{130, 205, 83, 7070, 10000, 123446}))
}
