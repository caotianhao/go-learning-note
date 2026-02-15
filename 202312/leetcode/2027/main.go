package main

import "fmt"

func minimumMoves(s string) (cnt int) {
	l := len(s)
	for i := 0; i < l; {
		if s[i] == 'O' {
			i++
		} else {
			i += 3
			cnt++
		}
	}
	return
}

func main() {
	fmt.Println(minimumMoves("XXOXOXOXOOXOXOXOX"))
	fmt.Println(minimumMoves("XXX"))  //1
	fmt.Println(minimumMoves("XXOX")) //2
	fmt.Println(minimumMoves("OOOO")) //0
}
