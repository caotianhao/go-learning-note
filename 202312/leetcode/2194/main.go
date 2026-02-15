package main

import "fmt"

func cellsInRange(s string) []string {
	charSlice, numSlice, targetSlice := make([]string, 0), make([]int, 0), make([]string, 0)
	for i := s[0]; i <= s[3]; i++ {
		charSlice = append(charSlice, string(i))
	}
	for i := s[1]; i <= s[4]; i++ {
		numSlice = append(numSlice, int(i)-48)
	}
	for i := 0; i < len(charSlice); i++ {
		for j := 0; j < len(numSlice); j++ {
			targetSlice = append(targetSlice, charSlice[i]+fmt.Sprintf("%d", numSlice[j]))
		}
	}
	return targetSlice
}

func main() {
	fmt.Println(cellsInRange("A2:E6"))
}
