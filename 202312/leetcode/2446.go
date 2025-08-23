package main

import (
	"fmt"
	"strconv"
)

func haveConflict(event1 []string, event2 []string) bool {
	b11, _ := strconv.ParseInt(event1[0][:2], 10, 64)
	b12, _ := strconv.ParseInt(event1[0][3:], 10, 64)
	e11, _ := strconv.ParseInt(event1[1][:2], 10, 64)
	e12, _ := strconv.ParseInt(event1[1][3:], 10, 64)
	begin1 := int(b11)*60 + int(b12)
	end1 := int(e11)*60 + int(e12)
	b21, _ := strconv.ParseInt(event2[0][:2], 10, 64)
	b22, _ := strconv.ParseInt(event2[0][3:], 10, 64)
	e21, _ := strconv.ParseInt(event2[1][:2], 10, 64)
	e22, _ := strconv.ParseInt(event2[1][3:], 10, 64)
	begin2 := int(b21)*60 + int(b22)
	end2 := int(e21)*60 + int(e22)
	return begin2 <= end1 && begin1 <= end2
}

func main() {
	fmt.Println(haveConflict([]string{"01:15", "02:00"}, []string{"03:15", "07:00"}))
}
