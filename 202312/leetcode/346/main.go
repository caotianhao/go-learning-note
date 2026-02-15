package main

import "fmt"

type MovingAverage346 struct {
	arr  []int
	size int
}

func Constructor346(size int) MovingAverage346 {
	return MovingAverage346{[]int{}, size}
}

func (ma *MovingAverage346) Next(val int) float64 {
	ma.arr = append(ma.arr, val)
	l := len(ma.arr)
	sum := 0
	if l <= ma.size {
		for _, v := range ma.arr {
			sum += v
		}
		return float64(sum) / float64(l)
	} else {
		for i := l - 1; i >= l-ma.size; i-- {
			sum += ma.arr[i]
		}
		return float64(sum) / float64(ma.size)
	}
}

func main() {
	a := Constructor346(3)
	fmt.Println(a.Next(1))
	fmt.Println(a.Next(10))
	fmt.Println(a.Next(3))
	fmt.Println(a.Next(5))
}
