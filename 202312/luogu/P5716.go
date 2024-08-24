package main

import "fmt"

func isRun(a int) bool {
	if a%4 == 0 && a%100 != 0 {
		return true
	}
	if a%400 == 0 {
		return true
	}
	return false
}

func main() {
	var y, m int
	_, _ = fmt.Scanf("%d %d", &y, &m)
	if m == 1 || m == 3 || m == 5 || m == 7 || m == 8 || m == 10 || m == 12 {
		fmt.Printf("%d", 31)
	} else if m == 4 || m == 6 || m == 9 || m == 11 {
		fmt.Printf("%d", 30)
	} else {
		if isRun(y) {
			fmt.Printf("%d", 29)
		} else {
			fmt.Printf("%d", 28)
		}
	}
}
