package main

import "fmt"

//func twoSum1(nums []int, target int) []int {
//	a, b := 0, 0
//	var my []int
//	my = make([]int, 0)
//loop:
//	for i := 0; i < len(nums); i++ {
//		for j := i + 1; j < len(nums); j++ {
//			if target == nums[i]+nums[j] {
//				a = i
//				b = j
//				break loop
//			}
//		}
//	}
//	my = append(my, a, b)
//	return my
//}

func twoSum1(nums []int, target int) (ret []int) {
	map1 := map[int]int{}
	for i, v := range nums {
		map1[v] = i
	}
	for i, v := range nums {
		if ind, ok := map1[target-v]; ok && i != ind {
			ret = append(ret, i)
			ret = append(ret, ind)
			return
		}
	}
	return
}

func main() {
	fmt.Println(twoSum1([]int{2, 7, 11, 15}, 17))
	fmt.Println(twoSum1([]int{3, 3}, 6))
}
