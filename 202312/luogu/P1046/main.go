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
	arr, cnt := make([]int, 0), 0
	scanner.Scan()
	t := scanner.Text()
	ta := strings.Split(t, " ")
	for _, v := range ta {
		temp, _ := strconv.Atoi(v)
		arr = append(arr, temp)
	}
	scanner.Scan()
	t = scanner.Text()
	temp, _ := strconv.Atoi(t)
	arr = append(arr, temp)
	for i := 0; i < 10; i++ {
		if arr[10]+30 >= arr[i] {
			cnt++
		}
	}
	fmt.Println(cnt)
}
