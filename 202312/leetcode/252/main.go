package main

import (
	"fmt"
	"sort"
)

func canAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][1] > intervals[i+1][0] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canAttendMeetings([][]int{{2, 4}, {7, 10}}))
}
