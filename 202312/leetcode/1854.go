package main

import "fmt"

func maximumPopulation(logs [][]int) int {
	mapPeople, l, maxN, maxYear := make([]int, 2055), len(logs), 0, 0
	for i := 0; i < l; i++ {
		for j := logs[i][0]; j < logs[i][1]; j++ {
			mapPeople[j]++
		}
	}
	for i, v := range mapPeople {
		if v > maxN {
			maxN = v
			maxYear = i
		}
	}
	return maxYear
}

func main() {
	fmt.Println(maximumPopulation([][]int{{1950, 1964}, {1960, 1971}, {1966, 1981}}))
	fmt.Println(maximumPopulation([][]int{{1993, 1999}, {2000, 2010}}))
}
