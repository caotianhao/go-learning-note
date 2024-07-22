package main

import (
	"fmt"
	"sort"
)

func main() {
	//随便写点 key-value
	testMap := map[int]int{}
	num := []int{1, 13, 7, 9, 27, 30, 4, 2, 12, 18, 49}
	testMap[num[0]] = 80
	testMap[num[1]] = 200
	testMap[num[2]] = 14
	testMap[num[3]] = 33
	testMap[num[4]] = 94
	testMap[num[5]] = 4
	testMap[num[6]] = 110
	testMap[num[7]] = 11
	testMap[num[8]] = 1
	testMap[num[9]] = 6
	testMap[num[10]] = 1104

	//sort by key
	//把 key 传到一个数组里，单独排序，然后根据排序后的 key 导出相应的 value
	//没啥意义
	key := make([]int, 0)
	for i := range testMap {
		key = append(key, i)
	}
	sort.Ints(key)
	fmt.Println("sort by key:")
	for i := range key {
		fmt.Printf("testMap[%d] = %d\n", key[i], testMap[key[i]])
	}

	//sort by value
	//map 可以根据 key 导出相应的 value，但不能根据 value 导出 key
	//因为 key 是唯一的，而 value 不是唯一的
	fmt.Println("-----------------------------------------")
	fmt.Println("sort by value:")
	//sort.Slice 有两个参数，第一个是待排序的切片
	//第二个是排序规则
	sort.Slice(num, func(a, b int) bool {
		//这里是根据 value 从大到小排列 key
		return testMap[num[a]] < testMap[num[b]]
		//这里是根据从大到小排列 key,效果就和 sort.Ints() 的反转效果一样
		//return num[a] > num[b]
	})
	for i := range num {
		fmt.Printf("testMap[%d] = %d\n", num[i], testMap[num[i]])
	}
	//fmt.Println(num)
}
