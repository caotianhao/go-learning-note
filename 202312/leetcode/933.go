package main

import "fmt"

type RecentCounter struct {
	recentCounter []int
}

func Constructor933() RecentCounter {
	return RecentCounter{[]int{}}
}

func (rc *RecentCounter) Ping(t int) int {
	cnt := 0
	rc.recentCounter = append(rc.recentCounter, t)
	for _, v := range rc.recentCounter {
		if v >= t-3000 && v <= t {
			cnt++
		}
	}
	return cnt
}

func main() {
	rc := Constructor933()
	fmt.Println(rc.Ping(1))
	fmt.Println(rc.Ping(100))
	fmt.Println(rc.Ping(3001))
	fmt.Println(rc.Ping(3002))
}
