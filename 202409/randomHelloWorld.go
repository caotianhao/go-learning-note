package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	dictionary string
	target     = "Hello World!"
)

func randomLetter() byte {
	rand.NewSource(time.Now().UnixNano())
	return dictionary[rand.Intn(len(dictionary))]
}

func main() {
	for i := 32; i <= 126; i++ {
		dictionary += string(rune(i))
	}
	result := make([]byte, len(target))
	currentIndex := 0
	cnt := 0

	for currentIndex < len(target) {
		letter := randomLetter()
		cnt++

		result[currentIndex] = letter
		fmt.Printf("%s\n", string(result[:currentIndex+1]))

		if letter == target[currentIndex] {
			currentIndex++
		}

		time.Sleep(30 * time.Millisecond)
	}
	fmt.Println(cnt)
}
