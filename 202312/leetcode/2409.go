package main

import (
	"fmt"
	"strconv"
)

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	days := [13]int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334, 365}
	tmp, _ := strconv.ParseInt(arriveAlice[:2], 10, 64)
	tmp1, _ := strconv.ParseInt(arriveAlice[3:], 10, 64)
	a1 := days[int(tmp)-1] + int(tmp1)
	tmp, _ = strconv.ParseInt(leaveAlice[:2], 10, 64)
	tmp1, _ = strconv.ParseInt(leaveAlice[3:], 10, 64)
	a2 := days[int(tmp)-1] + int(tmp1)
	tmp, _ = strconv.ParseInt(arriveBob[:2], 10, 64)
	tmp1, _ = strconv.ParseInt(arriveBob[3:], 10, 64)
	b1 := days[int(tmp)-1] + int(tmp1)
	tmp, _ = strconv.ParseInt(leaveBob[:2], 10, 64)
	tmp1, _ = strconv.ParseInt(leaveBob[3:], 10, 64)
	b2 := days[int(tmp)-1] + int(tmp1)
	fmt.Println(a1, a2, b1, b2)
	// alice 先
	if a1 <= b1 && b1 <= a2 {
		return min(a2-b1+1, min(a2-a1+1, b2-b1+1))
	}
	// bob 先
	if b1 <= a1 && a1 <= b2 {
		return min(b2-a1+1, min(a2-a1+1, b2-b1+1))
	}
	return 0
}

func main() {
	fmt.Println(countDaysTogether("08-15", "08-18", "08-16", "08-19")) // 3
	fmt.Println(countDaysTogether("10-01", "10-31", "11-01", "12-31")) // 0
	fmt.Println(countDaysTogether("01-25", "02-02", "01-28", "02-04")) // 6
	fmt.Println(countDaysTogether("09-01", "10-19", "06-19", "10-20")) // 49
}
