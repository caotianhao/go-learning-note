package main

import "fmt"

func slowestKey(releaseTimes []int, keys string) byte {
	timeKey, n := make([]int, 26), len(releaseTimes)
	part, maxTime := make([]int, n), -1
	part[0] = releaseTimes[0]
	for i := 1; i < n; i++ {
		part[i] = releaseTimes[i] - releaseTimes[i-1]
	}
	for i := 0; i < n; i++ {
		timeKey[int(keys[i]-'a')] = max(part[i], timeKey[int(keys[i]-'a')])
		maxTime = max(maxTime, timeKey[int(keys[i]-'a')])
	}
	for i := 25; i >= 0; i-- {
		if timeKey[i] == maxTime {
			return byte(i) + 'a'
		}
	}
	return '?'
}

func main() {
	fmt.Println(slowestKey([]int{9, 29, 49, 50}, "cbcd"))
}
