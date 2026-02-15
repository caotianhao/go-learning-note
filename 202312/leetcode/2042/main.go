package main

import (
	"fmt"
	"sort"
	"strconv"
)

func areNumbersAscending(s string) bool {
	tmp := divide2042(s)
	l, strSlice, numSlice := len(tmp), make([]string, 0), make([]int, 0)
	for i := 0; i < l; i++ {
		if tmp[i][0] >= '0' && tmp[i][0] <= '9' {
			strSlice = append(strSlice, tmp[i])
		}
	}
	for _, v := range strSlice {
		temp, _ := strconv.ParseInt(v, 10, 64)
		numSlice = append(numSlice, int(temp))
	}
	return sort.IntsAreSorted(numSlice) && !isHaveRedundant2042(numSlice)
}

func divide2042(s string) (slice2042 []string) {
	str := ""
	for _, v := range s {
		if v == ' ' {
			slice2042 = append(slice2042, str)
			str = ""
			continue
		}
		str += string(v)
	}
	slice2042 = append(slice2042, str)
	return
}

func isHaveRedundant2042(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(areNumbersAscending("1 box has 3 blue 4 red 6 green and 12 yellow marbles"))
	fmt.Println(areNumbersAscending("hello world 5 x 5"))
}
