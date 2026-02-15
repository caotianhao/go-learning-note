package main

import (
	"fmt"
	"sort"
)

//func sortPeople(names []string, heights []int) []string {
//	mySlice, myMap, l := make([]string, 0), make(map[int]string), len(heights)
//	for i := 0; i < l; i++ {
//		myMap[heights[i]] = names[i]
//	}
//	bubble(heights)
//	for i := 0; i < l; i++ {
//		mySlice = append(mySlice, myMap[heights[i]])
//	}
//	return mySlice
//}
//
//func bubble(arr []int) []int {
//	l := len(arr)
//	for i := 0; i < l; i++ {
//		for j := 0; j < l-i-1; j++ {
//			if arr[j] < arr[j+1] {
//				arr[j], arr[j+1] = arr[j+1], arr[j]
//			}
//		}
//	}
//	return arr
//}

func sortPeople(names []string, heights []int) (res []string) {
	m, l := map[int]string{}, len(names)
	for i := 0; i < l; i++ {
		m[heights[i]] = names[i]
	}
	sort.Sort(sort.Reverse(sort.IntSlice(heights)))
	for _, v := range heights {
		res = append(res, m[v])
	}
	return
}

func main() {
	arrString := []string{"Mary", "John", "Emma"}
	arrHeights := []int{180, 165, 170}
	fmt.Println(sortPeople(arrString, arrHeights))
}
