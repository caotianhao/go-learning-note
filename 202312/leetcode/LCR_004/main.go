package main

import "fmt"

func singleNumber2004(nums []int) int {
	var ans int32
	for i := 0; i < 32; i++ {
		var cnt int32
		for _, v := range nums {
			cnt += int32(v) >> i & 1
		}
		if cnt%3 == 1 {
			ans |= 1 << i
		}
	}
	return int(ans)
}

func main() {
	fmt.Println(singleNumber2004([]int{0, 1, 0, 1, 0, 1, 100}))
}
