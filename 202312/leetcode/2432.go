package main

import "fmt"

func hardestWorker(n int, logs [][]int) int {
	l := len(logs)
	for i := l - 1; i > 0; i-- {
		logs[i][1] = logs[i][1] - logs[i-1][1]
	}
	t, id := -1, n+3
	for _, v := range logs {
		if v[1] > t {
			t = v[1]
		}
	}
	for _, v := range logs {
		if v[1] == t {
			if id > v[0] {
				id = v[0]
			}
		}
	}
	return id
}

func main() {
	fmt.Println(hardestWorker(10, [][]int{{0, 3}, {2, 5}, {0, 9}, {1, 15}}))
}
