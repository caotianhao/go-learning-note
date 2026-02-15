package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func longestSubStr(arr []int) int {
	m := make(map[int]int)
	n := len(arr)
	if n == 0 {
		return 0
	}
	right, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, arr[i-1])
		}
		for right+1 < n && m[arr[right+1]] == 0 {
			m[arr[right+1]]++
			right++
		}
		ans = max(ans, right+1-i)
	}
	return ans
}

func main() {
	arr := make([]int, 0)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		data := strings.Split(input.Text(), " ")
		for _, v := range data {
			tmp, _ := strconv.Atoi(v)
			arr = append(arr, tmp)
		}
	}
	fmt.Println(longestSubStr(arr))
}
