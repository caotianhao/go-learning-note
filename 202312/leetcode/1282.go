package main

import "fmt"

func groupThePeople(groupSizes []int) (res [][]int) {
	groupMap := map[int][]int{}
	for i, v := range groupSizes {
		groupMap[v] = append(groupMap[v], i)
	}
	for i, v := range groupMap {
		tmp := make([]int, 0)
		for j := 0; j < len(v); j++ {
			tmp = append(tmp, v[j])
			if len(tmp)%i == 0 {
				res = append(res, tmp)
				tmp = make([]int, 0)
			}
		}
	}
	return
}

func main() {
	fmt.Println(groupThePeople([]int{3, 3, 3, 3, 3, 1, 3}))
	fmt.Println(groupThePeople([]int{1}))
}
