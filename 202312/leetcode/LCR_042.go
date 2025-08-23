package main

type RecentCounter2042 struct {
	recentCounter []int
}

func Constructor2042() RecentCounter2042 {
	return RecentCounter2042{[]int{}}
}

func (rc *RecentCounter2042) Ping(t int) int {
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
	c := Constructor2042()
	c.Ping(12)
}
