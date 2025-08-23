package main

import "fmt"

func maximum69Number(num int) int {
	if num == 6 || num == 9 {
		return 9
	}
	if num == 66 {
		return 96
	}
	if num == 69 || num == 96 || num == 99 {
		return 99
	}
	if num >= 600 && num < 700 {
		return num + 300
	}
	if num == 966 {
		return 996
	}
	if num == 969 || num == 996 || num == 999 {
		return 999
	}
	if num >= 6000 && num <= 7000 {
		return num + 3000
	}
	if num >= 9000 {
		return maximum69Number(num-9000) + 9000
	}
	return -1
}

func main() {
	fmt.Println(maximum69Number(999))
}
