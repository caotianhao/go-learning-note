package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	hashMap, dis := map[int]int{}, make([]int, 0)
	for _, v := range nums {
		hashMap[v]++
	}
	for i := 1; i <= len(nums); i++ {
		if hashMap[i] == 0 {
			dis = append(dis, i)
		}
	}
	return dis
}

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDisappearedNumbers([]int{1, 1}))
}
