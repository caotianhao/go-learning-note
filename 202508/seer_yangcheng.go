package main

import (
	"fmt"
	"sort"
)

const (
	bagCap     = 12
	colorRed   = "\033[31m"
	colorReset = "\033[0m"
)

var ycMap = make(map[int]int)

func fillAndCheckYcMap() []int {
	ycAll := [][]int{
		{3150, 3174, 3211, 3200, 3223, 2787, 3340, 2844, 2739, 3322, 3105, 2988}, // 1
		{3372, 3036, 2959, 3082, 4489, 4686, 4050, 4048, 4030, 4024, 3240, 3355}, // 2
		{3613, 3740, 3467, 3697, 3734, 4004, 3994, 3386, 3381, 3390, 3516, 4287}, // 3
		{3807, 3886, 3529, 4327, 3456, 3941, 4676, 4186, 4182, 3627, 4125, 4466}, // 4
		{3561, 4467, 3711, 4766, 4749, 4625, 4553, 4511, 4477, 4422, 4219, 4208}, // 5
		{4716, 4642, 4635, 4632, 4597, 4569, 4549, 4547, 4532, 4449, 4454, 4401}, // 6
		{4376, 4364, 4354, 4239, 4785, 4189, 4037, 4075, 3847, 3806, 3802, 3507}, // 7
		{4648, 4647, 5000, 4649, 3512, 3291, 4732, 4706, 4633, 4658, 4628, 4487},
		{4371, 4377, 4442, 4599, 4618, 4484, 4254, 4143, 3844, 3894, 3699, 3573},
		{2987, 3056, 3328, 3414, 3567, 3866, 4606, 4338, 4463, 3437, 3463, 3460},
		{4661, 3979, 3945, 4430, 3577, 4163, 3670, 4387, 3637, 4521, 4761, 4712}, // 8
		{3407, 3524, 3432, 3378, 4788, 4747, 3924, 4797, 4804, 4803, 4802, 4801}, // 9
	}

	for i := 0; i < len(ycAll); i++ {
		for j := 0; j < bagCap; j++ {
			ycMap[ycAll[i][j]]++
		}
	}

	if len(ycMap) != len(ycAll)*bagCap {
		res := make([]int, 0)
		for i, v := range ycMap {
			if v > 1 {
				res = append(res, i)
			}
		}
		return res
	}

	return nil
}

func listYcMap() {
	ycSlice := make([]int, 0)
	for item := range ycMap {
		ycSlice = append(ycSlice, item)
	}

	sort.Ints(ycSlice)
	for i := 0; i < len(ycSlice); i++ {
		if i%bagCap == 0 && i != 0 {
			fmt.Println()
		}
		fmt.Printf("%d\t", ycSlice[i])
	}

	fmt.Println()
}

func main() {
	res := fillAndCheckYcMap()
	if len(res) != 0 {
		fmt.Println(colorRed + "**********************************************" + colorReset)
		fmt.Printf(colorRed+"duplicate key: %v"+colorReset+"\n", res)
		fmt.Println(colorRed + "**********************************************" + colorReset)
	} else {
		listYcMap()
	}
}
