package main

import "fmt"

func furthestDistanceFromOrigin(moves string) int {
	l, r, z := 0, 0, 0
	for _, v := range moves {
		if v == 'L' {
			l++
		} else if v == 'R' {
			r++
		} else {
			z++
		}
	}
	return z + abs2833(l-r)
}

func abs2833(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func main() {
	fmt.Println(furthestDistanceFromOrigin("L_RL__R"))
}
