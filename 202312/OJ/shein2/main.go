package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func threeNumbersSum(arr []int) (ans [][]int) {
	sort.Ints(arr)
	l := len(arr)
	hm := map[int]struct{}{}
	for i, number := range arr[:l-2] {
		if i > 0 && number == arr[i-1] {
			continue
		}
		if number+arr[i+1]+arr[i+2] > 0 {
			break
		}
		if number+arr[l-2]+arr[l-1] < 0 {
			continue
		}
		idx1, idx2 := i+1, l-1
		for idx1 < idx2 {
			tempSum := number + arr[idx1] + arr[idx2]
			if tempSum > 0 {
				idx2--
			} else if tempSum < 0 {
				idx1++
			} else {
				temp := []int{number, arr[idx1], arr[idx2]}
				_, ok1 := hm[number]
				_, ok2 := hm[arr[idx1]]
				_, ok3 := hm[arr[idx2]]
				if !ok1 && !ok2 && !ok3 {
					hm[number] = struct{}{}
					hm[arr[idx1]] = struct{}{}
					hm[arr[idx2]] = struct{}{}
					sort.Ints(temp)
					ans = append(ans, temp)
					for idx1++; idx1 < idx2 && arr[idx1] == arr[idx1-1]; idx1++ {
					}
					for idx2--; idx1 < idx2 && arr[idx2] == arr[idx2+1]; idx2-- {
					}
				} else {
				}
			}
		}
	}
	return
}

func main() {
	arr := make([]int, 0)
	input := bufio.NewReader(os.Stdin)
	r, _ := input.ReadString('\n')
	r = strings.Trim(r, "\n")
	data := strings.Split(r, " ")
	for _, v := range data {
		t, _ := strconv.Atoi(v)
		arr = append(arr, t)
	}
	fmt.Println(threeNumbersSum(arr))
}
