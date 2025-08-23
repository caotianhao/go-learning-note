package main

import "fmt"

func intersect(nums1 []int, nums2 []int) (res []int) {
	m1, m2 := map[int]int{}, map[int]int{}
	for _, v := range nums1 {
		m1[v]++
	}
	for _, v := range nums2 {
		m2[v]++
	}
	for i, v := range m1 {
		if v2, ok := m2[i]; ok {
			for j := 0; j < min(v, v2); j++ {
				res = append(res, i)
			}
		}
	}
	return
}

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
}
