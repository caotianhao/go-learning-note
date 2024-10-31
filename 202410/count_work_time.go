package main

import (
	"fmt"
	"time"
)

var (
	layout       = time.TimeOnly
	totalMinute  = 0
	targetMinute = 9 * 60
)

func main() {
	everyDayTime := [][2]string{
		{"09:30:03", "18:11:09"},
		{"09:40:19", "20:54:23"},
		{"09:35:59", "18:40:23"},
		//{"09:00:00", "18:00:00"},
		//{"09:00:00", "18:00:00"},
		//{"09:00:00", "18:00:00"},
		//{"09:00:00", "18:00:00"},
	}

	for cnt, times := range everyDayTime {
		startTime, _ := time.Parse(layout, times[0])
		endTime, _ := time.Parse(layout, times[1])

		diff := endTime.Sub(startTime)
		workMinute := int(diff.Minutes())
		minuteDiff := int(diff.Minutes()) - targetMinute
		totalMinute += minuteDiff

		fmt.Printf("Day %d: 工作时间 %d 分钟，差距 %d 分钟\n", cnt+1, workMinute, minuteDiff)
	}

	fmt.Println(totalMinute)
}
