package main

import "fmt"

func processQueries(queries []int, m int) (res []int) {
	p := make([]int, m)
	for i := 0; i < m; i++ {
		p[i] = i + 1
	}
	for i := 0; i < len(queries); i++ {
		for j := 0; j < m; j++ {
			if queries[i] == p[j] {
				res = append(res, j)
				p = change1409(p, j)
			}
		}
	}
	return
}

func change1409(arr []int, ind int) []int {
	res := make([]int, 1)
	res[0] = arr[ind]
	res = append(res, arr[:ind]...)
	res = append(res, arr[ind+1:]...)
	return res
}

func main() {
	fmt.Println(processQueries([]int{8, 5, 5, 8, 8}, 8)) // 6 5 0 7 5
}
