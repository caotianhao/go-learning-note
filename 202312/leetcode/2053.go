package main

import "fmt"

func kthDistinct(arr []string, k int) string {
	mapKth, indexSlice := map[string]int{}, make([]int, 0)
	for _, v := range arr {
		mapKth[v]++
	}
	for i, v := range arr {
		if mapKth[v] == 1 {
			indexSlice = append(indexSlice, i)
		}
	}
	if k <= len(indexSlice) {
		return arr[indexSlice[k-1]]
	} else {
		return ""
	}
}

func main() {
	fmt.Println(kthDistinct([]string{"d", "b", "c", "b", "c", "a"}, 2))
	fmt.Println(kthDistinct([]string{"aaa", "aa", "a"}, 1))
	fmt.Println(kthDistinct([]string{"a", "b", "a"}, 3))
}
