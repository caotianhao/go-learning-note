package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func minCostTickets(days []int, costs []int) int {
	year := make([]int, days[len(days)-1]+1)
	for k, v := range days {
		if k != 0 {
			for i := days[k-1] + 1; i < v; i++ {
				year[i] = year[days[k-1]]
			}
		}
		temp1 := min222(costs[0]+year[max2222(v-1, 0)], costs[1]+year[max2222(v-7, 0)])
		year[v] = min222(temp1, costs[2]+year[max2222(v-30, 0)])
	}
	return year[days[len(days)-1]]
}

func min222(x int, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func max2222(i int, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func main() {
	arr := make([]int, 0)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	t := sc.Text()
	f := strings.Fields(t)
	for _, v := range f {
		t, _ := strconv.Atoi(v)
		arr = append(arr, t)
	}

	ccc := make([]int, 0)
	sc.Scan()
	t = sc.Text()
	f = strings.Fields(t)
	for _, v := range f {
		t, _ := strconv.Atoi(v)
		ccc = append(ccc, t)
	}

	fmt.Println(arr, ccc)
	fmt.Println(minCostTickets(arr, ccc))
}
