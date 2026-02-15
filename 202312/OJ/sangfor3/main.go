package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func longestDeArr(arr []int) int {
	n := len(arr)
	ans := -1
	if n == 1 {
		return 1
	}
	left := 0
	for i, v := range arr {
		if i > 0 && v >= arr[i-1] {
			left = i
		}
		ans = max(ans, i-left+1)
	}
	return ans
}

func main() {
	arr := make([]int, 0)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		data := strings.Split(input.Text(), ",")
		for _, v := range data {
			tmp, _ := strconv.Atoi(v)
			arr = append(arr, tmp)
		}
	}
	fmt.Println(longestDeArr(arr))
}
