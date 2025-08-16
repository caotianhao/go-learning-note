package main

import (
	"fmt"
	"sort"
)

var ycMap = make(map[int]struct{})

func increaseYcMap() {
	ycAll := [][]int{
		{3150, 3174, 3211, 3200, 3223, 2787, 3340, 2844, 2739, 3322, 3105, 2988}, // 1
		{3372, 3036, 2959, 3082, 4489, 4686, 4050, 4048, 4030, 4024, 3240, 3355}, // 2
		{3613, 3740, 3467, 3697, 3734, 4004, 3994, 3386, 3381, 3390, 3516, 4287}, // 3
		{3807, 3886, 3529, 4327, 3456, 3941, 4676, 4186, 4182, 3627, 4125, 4466}, // 4
		{3561, 4467, 3711, 4766, 4749, 4625, 4553, 4511, 4477, 4422, 4219, 4208}, // 5
	}

	for i := 0; i < len(ycAll); i++ {
		for j := 0; j < len(ycAll[i]); j++ {
			ycMap[ycAll[i][j]] = struct{}{}
		}
	}

	fmt.Printf("【len = %d】\n", len(ycMap))
}

func listYcMap() {
	ycSlice := make([]int, 0)
	for item := range ycMap {
		ycSlice = append(ycSlice, item)
	}

	sort.Ints(ycSlice)
	for i := 0; i < len(ycSlice); i++ {
		if i%10 == 0 && i != 0 {
			fmt.Println()
		}
		fmt.Printf("%d\t", ycSlice[i])
	}
}

func main() {
	increaseYcMap()
	listYcMap()
}
