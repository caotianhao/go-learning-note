package main

import "fmt"

func findingUsersActiveMinutes(logs [][]int, k int) []int {
	logMap := map[int]map[int]struct{}{}
	for _, v := range logs {
		logMap[v[0]] = map[int]struct{}{}
	}
	for _, v := range logs {
		logMap[v[0]][v[1]] = struct{}{}
	}
	res := make([]int, k)
	for _, v := range logMap {
		res[len(v)-1]++
	}
	return res
}

func main() {
	fmt.Println(findingUsersActiveMinutes([][]int{{0, 5}, {1, 2}, {0, 2}, {0, 5}, {1, 3}}, 5))
}
