package main

import "fmt"

const (
	NumberMin = 1
	NumberMax = 5000
)

func main() {
	invalidNumbers := []int{4399}
	fmt.Printf("have %d\nlose:", NumberMax-NumberMin+1)

	for i := 0; i < len(invalidNumbers); i++ {
		if i%10 == 0 {
			fmt.Println()
		}
		fmt.Printf("%d ", invalidNumbers[i])
	}
	
	fmt.Println()
}
