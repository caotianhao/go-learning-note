package main

import "fmt"

func finalValueAfterOperations(operations []string) int {
	a := 0
	for i := 0; i < len(operations); i++ {
		if operations[i] == "X++" || operations[i] == "++X" {
			a++
		} else {
			a--
		}
	}
	return a
}

func main() {
	arr := []string{"X++", "++X", "--X", "X--"}
	fmt.Println(finalValueAfterOperations(arr))
}
