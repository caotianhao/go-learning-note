package main

import "fmt"

func findTheArrayConcVal(nums []int) int64 {
	l := len(nums)
	var answer int64 = 0
	for i := 0; i < l/2; i++ {
		a, b := nums[i], nums[l-1-i]
		if b >= 1 && b <= 9 {
			answer += int64(a*10 + b)
		} else if b >= 10 && b <= 99 {
			answer += int64(a*100 + b)
		} else if b >= 100 && b <= 999 {
			answer += int64(a*1000 + b)
		} else if b >= 1000 && b <= 9999 {
			answer += int64(a*10000 + b)
		} else {
			answer += int64(a*100000 + b)
		}
	}
	if l%2 == 1 {
		answer += int64(nums[l/2])
	}
	return answer
}

func main() {
	fmt.Println(findTheArrayConcVal([]int{7, 52, 2, 4}))      //596
	fmt.Println(findTheArrayConcVal([]int{5, 14, 13, 8, 12})) //673
	fmt.Println(findTheArrayConcVal([]int{2}))                //2
	fmt.Println(findTheArrayConcVal([]int{5, 14, 13, 8, 12, 15, 10000, 354, 24}))
}
