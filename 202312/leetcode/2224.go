package main

import (
	"fmt"
	"strconv"
)

func convertTime(current string, correct string) (cnt int) {
	tmp1, _ := strconv.ParseInt(current[:2], 10, 64)
	tmp2, _ := strconv.ParseInt(current[3:], 10, 64)
	begin := 60*int(tmp1) + int(tmp2)
	tmp1, _ = strconv.ParseInt(correct[:2], 10, 64)
	tmp2, _ = strconv.ParseInt(correct[3:], 10, 64)
	end := 60*int(tmp1) + int(tmp2)
	need := end - begin
	for need != 0 {
		if need >= 60 {
			need -= 60
			cnt++
		} else if need < 60 && need >= 15 {
			need -= 15
			cnt++
		} else if need < 15 && need >= 5 {
			need -= 5
			cnt++
		} else {
			need -= 1
			cnt++
		}
	}
	return
}

func main() {
	fmt.Println(convertTime("02:30", "04:35"))
	fmt.Println(convertTime("00:00", "23:59"))
}
