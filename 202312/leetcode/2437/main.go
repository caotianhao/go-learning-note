package main

import "fmt"

func countTime(time string) int {
	ans := 1
	s1, s2 := time[:2], time[3:]
	if s1 == "0?" || s1 == "1?" {
		ans *= 10
	} else if s1 == "2?" {
		ans *= 4
	} else if s1 == "??" {
		ans *= 24
	} else if s1 == "?0" || s1 == "?1" ||
		s1 == "?2" || s1 == "?3" {
		ans *= 3
	} else if s1 == "?4" || s1 == "?5" || s1 == "?6" ||
		s1 == "?7" || s1 == "?8" || s1 == "?9" {
		ans *= 2
	}
	if s2 == "??" {
		ans *= 60
	} else if s2 == "?0" || s2 == "?1" || s2 == "?2" ||
		s2 == "?3" || s2 == "?4" || s2 == "?5" ||
		s2 == "?6" || s2 == "?7" || s2 == "?8" || s2 == "?9" {
		ans *= 6
	} else if s2 == "0?" || s2 == "1?" || s2 == "2?" ||
		s2 == "3?" || s2 == "4?" || s2 == "5?" {
		ans *= 10
	}
	return ans
}

func main() {
	fmt.Println(countTime("0?:22"))
	fmt.Println(countTime("2?:??"))
}
