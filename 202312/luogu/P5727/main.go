package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	d, _ := strconv.Atoi(scanner.Text())
	arr := make([]int, 0)
	for d != 1 {
		arr = append(arr, d)
		if d&1 == 0 {
			d /= 2
		} else {
			d = d*3 + 1
		}
	}
	arr = append(arr, 1)
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Printf("%d ", arr[i])
	}
}
