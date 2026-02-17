package main

import (
	"fmt"
	"strings"
)

func shortestCompletingWord(licensePlate string, words []string) string {
	needDelete := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", " "}
	for i := 0; i < 11; i++ {
		licensePlate = strings.Replace(licensePlate, needDelete[i], "", -1)
	}
	licensePlate = strings.ToLower(licensePlate)
	lpMap, trueInd, minN := map[string]int{}, 0, 16
	for i := range licensePlate {
		lpMap[string(licensePlate[i])]++
	}
	for i := range words {
		tmp := map[string]int{}
		for j := 0; j < len(words[i]); j++ {
			tmp[string(words[i][j])]++
		}
		yes := true
		for ind, val := range lpMap {
			tmpVal, ok := tmp[ind]
			if !ok || tmpVal < val {
				yes = false
			}
		}
		if yes {
			if len(words[i]) < minN {
				minN = len(words[i])
				trueInd = i
			}
		}
	}
	return words[trueInd]
}

func main() {
	fmt.Println(shortestCompletingWord("1s3 PSt", []string{"step", "steps", "stripe", "stepple"}))
	fmt.Println(shortestCompletingWord("1s3 456", []string{"looks", "pest", "stew", "show"}))
	fmt.Println(shortestCompletingWord("1s3 456", []string{"lo2os", "pes3t", "stew", "show"}))
}
