package main

import "fmt"

//注意原题要求要修改数组，不仅仅是返回一个int而已
func removeElement(nums []int, val int) int {
	l, cnt := len(nums), 0
	for i := 0; i < l; i++ {
		if val != nums[i] {
			nums[cnt] = nums[i]
			cnt++
		}
	}
	return cnt
}

func main() {
	arr := []int{3, 2, 2, 3}
	fmt.Println(removeElement(arr, 3))
}
