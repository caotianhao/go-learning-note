package main

import "fmt"

func attackTieBa(yes, no []int) int {
	l := len(yes)
	big, small := 0, 0
	for i := 0; i < l; i++ {
		if yes[i] > no[i] {
			big += yes[i] - no[i]
		} else if yes[i] < no[i] {
			small += no[i] - yes[i]
		}
	}
	if big > small {
		return big
	}
	return small
}

func main() {
	n := 0
	_, _ = fmt.Scanf("%d", &n)
	yes, no, yesList, noList := 0, 0, make([]int, 0), make([]int, 0)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &yes)
		yesList = append(yesList, yes)
	}
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &no)
		noList = append(noList, no)
	}
	fmt.Println(attackTieBa(yesList, noList))
}
