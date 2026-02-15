package main

import "fmt"

//func twoSum167(numbers []int, target int) []int {
//	hashMap := map[int]int{}
//	for i, v := range numbers {
//		if pos, ok := hashMap[target-v]; ok {
//			if i > pos {
//				return []int{pos + 1, i + 1}
//			}
//			return []int{i + 1, pos + 1}
//		}
//		hashMap[v] = i
//	}
//	return nil
//}

func twoSum167(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l <= r {
		if numbers[l]+numbers[r] < target {
			l++
		} else if numbers[l]+numbers[r] == target {
			return []int{l + 1, r + 1}
		} else {
			r--
		}
	}
	return []int{-1, -1}
}

func main() {
	//fmt.Println(twoSum167([]int{1, 2, 4, 4, 10}, 8))
	//fmt.Println(twoSum167([]int{2, 7, 11, 15}, 9))
	//fmt.Println(twoSum167([]int{2, 3, 4}, 6))
	//fmt.Println(twoSum167([]int{-1, 0}, -1))
	fmt.Println(twoSum167([]int{5, 25, 75}, 100))
}
