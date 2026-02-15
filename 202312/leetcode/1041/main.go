package main

import "fmt"

func isRobotBounded(instructions string) bool {
	direction := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	curDirection := 0 // 0北 1东 2南 3西
	x, y := 0, 0
	for _, v := range instructions {
		if v == 'G' {
			x += direction[curDirection][0]
			y += direction[curDirection][1]
		} else if v == 'L' {
			curDirection += 3
			curDirection %= 4
		} else {
			curDirection++
			curDirection %= 4
		}
	}
	// 不在原点 且 方向为北
	// 题目要求陷入循环为 true
	return curDirection != 0 || x == 0 && y == 0
}

func main() {
	fmt.Println(isRobotBounded("GGLLGG"))    //t
	fmt.Println(isRobotBounded("GG"))        //f
	fmt.Println(isRobotBounded("GL"))        //t
	fmt.Println(isRobotBounded("GLGLGGLGL")) //f
}
