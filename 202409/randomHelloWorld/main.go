package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var (
	dictionary string
	target     = "September 29, 2024, in Suwon, South Korea"
)

func randomLetter() byte {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return dictionary[r.Intn(len(dictionary))]
}

func main() {
	for i := 32; i <= 126; i++ {
		dictionary += string(rune(i))
	}

	result := make([]byte, len(target))
	currentIndex := 0
	currentPart := map[string]int{}
	cnt := 0

	for currentIndex < len(target) {
		letter := randomLetter()
		cnt++

		result[currentIndex] = letter
		partResult := string(result[:currentIndex+1])

		if letter == target[currentIndex] {
			currentIndex++
			currentPart[partResult] = cnt
		}
	}

	keys := make([]string, 0, len(currentPart))
	for k := range currentPart {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return currentPart[keys[i]] < currentPart[keys[j]]
	})

	fmt.Printf("%-50s %10s\n", "Part", "Count")
	fmt.Println()
	for _, k := range keys {
		fmt.Printf("%-50s %10d\n", k, currentPart[k])
	}
}
