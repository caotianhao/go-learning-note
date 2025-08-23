package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	hashMap := map[int]int{}
	for i, v := range numbers {
		if pos, ok := hashMap[target-v]; ok {
			if i > pos {
				return []int{pos, i}
			}
			return []int{i, pos}
		}
		hashMap[v] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{1, 2, 4, 4, 10}, 8))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
