package main

import "fmt"

func numberOfLines(widths []int, s string) []int {
	row, exSpace := 1, 0
	for _, v := range s {
		if exSpace+widths[int(v)-97] > 100 {
			exSpace = 0
			exSpace += widths[int(v)-97]
			row++
		} else {
			exSpace += widths[int(v)-97]
		}
	}
	return []int{row, exSpace}
}

func main() {
	fmt.Println(numberOfLines([]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
		10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(numberOfLines([]int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
		10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "bbbcccdddaaa"))
	fmt.Println(numberOfLines([]int{4, 2, 7, 6, 4, 2, 10, 2, 4, 6, 3, 6, 9, 10, 7,
		10, 10, 5, 10, 10, 8, 10, 10, 10, 10, 10}, "qfkjhafkjshufhkjhsdkjgfkaffadbfdsfabfgs"))
}
