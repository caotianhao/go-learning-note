package main

import "fmt"

func judgeCircle(moves string) bool {
	l, sx, zy := len(moves), 0, 0
	for i := 0; i < l; i++ {
		if moves[i] == 'U' {
			sx++
		} else if moves[i] == 'D' {
			sx--
		} else if moves[i] == 'L' {
			zy--
		} else {
			zy++
		}
	}
	if sx == 0 && zy == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(judgeCircle("DULLDUDL"))
}
