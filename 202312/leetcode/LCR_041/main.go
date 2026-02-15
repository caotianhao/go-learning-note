package main

import "fmt"

type MovingAverage struct {
	arr []int
	avg int
}

func Constructor2041(size int) MovingAverage {
	slice2041 := make([]int, 0)
	return MovingAverage{slice2041, size}
}

func (ma *MovingAverage) Next(val int) float64 {
	ma.arr = append(ma.arr, val)
	cnt, sum, l := 0, 0, len(ma.arr)
	if l <= ma.avg {
		for _, v := range ma.arr {
			sum += v
		}
		return float64(sum) / float64(l)
	}
	for i := l - 1; cnt < ma.avg; i-- {
		sum += ma.arr[i]
		cnt++
	}
	return float64(sum) / float64(ma.avg)
}

//输入：
//inputs = ["MovingAverage", "next", "next", "next", "next"]
//inputs = [    [3],          [1],    [10],   [3],    [5]  ]
//输出：
//[             null,          1.0,    5.5,  4.66667,  6.0]
func main() {
	a := Constructor2041(3)
	fmt.Println(a.Next(1))
	fmt.Println(a.Next(10))
	fmt.Println(a.Next(3))
	fmt.Println(a.Next(5))
}
