package main

import (
	"fmt"
	"strings"
)

func garbageCollection(garbage []string, travel []int) (timeGarbage int) {
	lastG, lastM, lastP, lg := -1, -1, -1, len(garbage)
	s, total := "", map[byte]int{}
	for i := 0; i < lg; i++ {
		s += garbage[i]
	}
	for i := 0; i < len(s); i++ {
		total[s[i]]++
	}
	for i := lg - 1; i >= 0; i-- {
		if strings.Contains(garbage[i], "G") {
			lastG = i
			break
		}
	}
	for i := lg - 1; i >= 0; i-- {
		if strings.Contains(garbage[i], "M") {
			lastM = i
			break
		}
	}
	for i := lg - 1; i >= 0; i-- {
		if strings.Contains(garbage[i], "P") {
			lastP = i
			break
		}
	}
	for i := 0; i < lastG; i++ {
		timeGarbage += travel[i]
	}
	for i := 0; i < lastM; i++ {
		timeGarbage += travel[i]
	}
	for i := 0; i < lastP; i++ {
		timeGarbage += travel[i]
	}
	return timeGarbage + total['G'] + total['M'] + total['P']
}

func main() {
	fmt.Println(garbageCollection([]string{"G", "P", "GP", "GG"}, []int{2, 4, 3}))
	fmt.Println(garbageCollection([]string{"MMM", "PGM", "GP"}, []int{3, 10}))
}
