package main

import "fmt"

func main() {
	var (
		n      int
		num    int64
		oriSum int64
		maxN   int64 = -1
	)
	_, _ = fmt.Scanf("%d", &n)
	arr := make([]int64, 0)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &num)
		arr = append(arr, num)
	}
	for _, v := range arr {
		oriSum += v
	}
	for i := 0; i < n-1; i++ {
		t := oriSum - arr[i] - arr[i+1]
		tt := t + arr[i]*arr[i+1]
		if tt > maxN {
			maxN = tt
		}
	}
	fmt.Println(maxN)
}
