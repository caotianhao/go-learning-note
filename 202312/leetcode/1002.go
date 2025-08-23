package main

import (
	"fmt"
	"math"
)

func commonChars(words []string) []string {
	l, slice1002 := len(words), make([]string, 0)
	charMap := map[string]int{}
	for i := 0; i < len(words[0]); i++ {
		charMap[string(words[0][i])]++
	}
	for i := 1; i < l; i++ {
		tempCharMap := map[string]int{}
		for j := 0; j < len(words[i]); j++ {
			tempCharMap[string(words[i][j])]++
		}
		for ind, val := range tempCharMap {
			if v, ok := charMap[ind]; ok {
				charMap[ind] = int(math.Min(float64(val), float64(v)))
			}
		}
		for ind := range charMap {
			if _, ok := tempCharMap[ind]; !ok {
				delete(charMap, ind)
			}
		}
	}
	for i, v := range charMap {
		for j := 0; j < v; j++ {
			slice1002 = append(slice1002, i)
		}
	}
	return slice1002
}

func main() {
	fmt.Println(commonChars([]string{"bella", "label", "roller"}))
	fmt.Println(commonChars([]string{"bella"}))
}
