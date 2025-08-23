package main

import "fmt"

type SparseVector struct {
	arr []int
}

func Constructor1570(nums []int) SparseVector {
	return SparseVector{nums}
}

// Return the dotProduct of two sparse vectors
func (s1 *SparseVector) dotProduct(s2 SparseVector) int {
	pro := 0
	for i := 0; i < len(s2.arr); i++ {
		pro += s1.arr[i] * s2.arr[i]
	}
	return pro
}

func main() {
	a1 := Constructor1570([]int{1, 0, 0, 2, 3})
	a2 := Constructor1570([]int{0, 3, 0, 4, 0})
	fmt.Println(a1.dotProduct(a2))
	fmt.Println(a2.dotProduct(a1))
}
