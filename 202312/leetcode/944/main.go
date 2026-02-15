package main

import (
	"fmt"
	"sort"
)

func minDeletionSize(strs []string) int {
	l, l1, col := len(strs), len(strs[0]), 0
	for j := 0; j < l1; j++ {
		slice944 := make([]int, 0)
		for i := 0; i < l; i++ {
			slice944 = append(slice944, int(strs[i][j]))
		}
		tmpSlice944 := make([]int, 0)
		for _, v := range slice944 {
			tmpSlice944 = append(tmpSlice944, v)
		}
		//fmt.Println("t", tmpSlice944)
		sort.Slice(slice944, func(i, j int) bool {
			return slice944[i] < slice944[j]
		})
		//fmt.Println("s", slice944)
		flag := true
		for a := 0; a < len(slice944); a++ {
			if tmpSlice944[a] != slice944[a] {
				flag = false
				break
			}
		}
		//fmt.Println(flag)
		if !flag {
			col++
		}
		//fmt.Println(i, col)
	}
	return col
}

func main() {
	fmt.Println(minDeletionSize([]string{"cba", "daf"}))
	fmt.Println(minDeletionSize([]string{"a", "b"}))
	fmt.Println(minDeletionSize([]string{"zyx", "wvu", "tsr"}))
	fmt.Println(minDeletionSize([]string{"zyx"}))
	fmt.Println(minDeletionSize([]string{"zyx", "cba", "daf", "ghi", "wvu", "tsr"}))
}
