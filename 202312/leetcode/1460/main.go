package main

import "fmt"

func canBeEqual(target []int, arr []int) bool {
	//思想：只要 target 和 arr 两数组中元素完全相同，个数也相同
	//也就是排序之后完全相同，即可在有限次翻转内使原 arr 变为 target
	targetMap, arrMap := make(map[int]int), make(map[int]int)
	lt, la, cnt := len(target), len(arr), 0
	for i := 0; i < lt; i++ {
		targetMap[target[i]]++
	}
	for i := 0; i < la; i++ {
		arrMap[arr[i]]++
	}
	for i, v := range targetMap {
		if v1, ok := arrMap[i]; ok && v == v1 {
			cnt++
		}
	}
	if cnt == len(targetMap) {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(canBeEqual([]int{1, 2, 3, 4}, []int{2, 4, 1, 3}))
}
