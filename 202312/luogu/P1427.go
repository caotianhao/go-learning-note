package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	d := strings.Split(scanner.Text(), " ")
	arr := make([]int, 0)
	for _, v := range d {
		t, _ := strconv.Atoi(v)
		arr = append(arr, t)
	}
	for i := len(arr) - 2; i >= 0; i-- {
		fmt.Printf("%d ", arr[i])
	}
}
