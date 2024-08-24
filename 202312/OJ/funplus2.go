package main

import (
	"fmt"
)

const (
	ep  = 1e-10
	add = 0
	sub = 1
	mul = 2
	div = 3
)

func judgePoints(cards []int, points int) bool {
	f64nums := make([]float64, 0)
	for _, v := range cards {
		f64nums = append(f64nums, float64(v))
	}
	var get24 func([]float64, int) bool
	get24 = func(arr []float64, target int) bool {
		if len(arr) == 0 {
			return false
		}
		if len(arr) == 1 {
			return absFp(arr[0]-float64(target)) < ep
		}
		l := len(arr)
		for i := 0; i < l; i++ {
			for j := 0; j < l; j++ {
				if i != j {
					tmp := make([]float64, 0)
					for k := 0; k < l; k++ {
						if i != k && j != k {
							tmp = append(tmp, arr[k])
						}
					}
					for flag := 0; flag < 4; flag++ {
						if (flag == 0 || flag == 2) && i < j {
							continue
						}
						switch flag {
						case add:
							tmp = append(tmp, arr[i]+arr[j])
						case sub:
							tmp = append(tmp, arr[i]-arr[j])
						case mul:
							tmp = append(tmp, arr[i]*arr[j])
						case div:
							if absFp(arr[j]) < ep {
								continue
							} else {
								tmp = append(tmp, arr[i]/arr[j])
							}
						}
						if get24(tmp, target) {
							return true
						}
						tmp = tmp[:len(tmp)-1]
					}
				}
			}
		}
		return false
	}
	return get24(f64nums, points)
}

func absFp(a float64) float64 {
	if a > 0 {
		return a
	}
	return -a
}

func main() {
	fmt.Println(judgePoints([]int{1, 2, 3, 4}, 24))
}
