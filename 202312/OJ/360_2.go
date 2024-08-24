package main

import "fmt"

func main() {
	steps := 0
	_, _ = fmt.Scanf("%d", &steps)
	x1, y1, x2, y2 := 0, 0, 0, 0
	_, _ = fmt.Scanf("%d %d %d %d", &x1, &y1, &x2, &y2)
	s1, s2 := "", ""
	_, _ = fmt.Scanf("%s", &s1)
	_, _ = fmt.Scanf("%s", &s2)
	for i := 0; i < steps; i++ {
		if x1 == x2 && y1 == y2 {
			fmt.Println(i)
			return
		}
		if s1[i] == 'W' {
			y1++
		} else if s1[i] == 'A' {
			x1--
		} else if s1[i] == 'S' {
			y1--
		} else {
			x1++
		}
		if s2[i] == 'W' {
			y2++
		} else if s2[i] == 'A' {
			x2--
		} else if s2[i] == 'S' {
			y2--
		} else {
			x2++
		}
	}
	if x1 == x2 && y1 == y2 {
		fmt.Println(steps)
	} else {
		fmt.Println("Not Over")
	}
}
