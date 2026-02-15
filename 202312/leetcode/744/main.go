package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
	for _, v := range letters {
		if v > target {
			return v
		}
	}
	return letters[0]
}

func main() {
	fmt.Print(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a'))
}
