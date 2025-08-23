package main

import (
	"fmt"
	"math/rand"
	"time"
)

func countTriplets982(nums []int) (cnt int) {
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			for k := 0; k < l; k++ {
				if nums[i]&nums[j]&nums[k] == 0 {
					cnt++
				}
			}
		}
	}
	return
}

func main() {
	fmt.Println(countTriplets982([]int{1, 2, 3}))
	fmt.Println(countTriplets982([]int{0, 0, 0}))
	rand.NewSource(time.Now().UnixNano())
	arr := make([]int, 0)
	for i := 0; i < 1000; i++ {
		arr = append(arr, rand.Intn(65536))
	}
	//fmt.Print("[")
	//for i := 0; i < 1000; i++ {
	//	fmt.Print(arr[i], ",")
	//}
	//fmt.Println("]")
	t := time.Now()
	countTriplets982(arr)
	t1 := time.Since(t)
	fmt.Println(t1)
}
