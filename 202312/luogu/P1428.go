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
	n, _ := strconv.Atoi(scanner.Text())
	arr, res := make([]int, 0), make([]int, 0)
	scanner.Scan()
	d := strings.Split(scanner.Text(), " ")
	for _, v := range d {
		t, _ := strconv.Atoi(v)
		arr = append(arr, t)
	}
	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				cnt++
			}
		}
		res = append(res, cnt)
	}
	for i := 0; i < len(res); i++ {
		fmt.Printf("%d ", res[i])
	}
}
