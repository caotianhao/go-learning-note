package main

import "fmt"

func interpret(command string) string {
	l, target := len(command), ""
	for i := 0; i < l; i++ {
		if command[i] == 'G' {
			target += "G"
		} else if command[i] == '(' {
			if command[i+1] == ')' {
				target += "o"
			} else {
				target += "al"
			}
		}
	}
	return target
}

func main() {
	a := "Goal"
	fmt.Println(interpret(a))
}
