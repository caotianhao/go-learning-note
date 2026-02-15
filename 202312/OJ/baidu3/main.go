package main

import (
	"fmt"
)

func findMinByDiv(arr *[]int) int {
	return div1(arr, 0, len(*arr)-1)
}

func div1(arr *[]int, left, right int) int {
	if left == right {
		return (*arr)[left]
	}
	middle := (left + right) / 2
	leftMin := div1(arr, left, middle)
	rightMin := div1(arr, middle+1, right)

	if leftMin < rightMin {
		return leftMin
	}
	return rightMin
}

func main() {
	n, m := 0, 0
	_, _ = fmt.Scanf("%d %d", &n, &m)
	angryMax := make([]int, n+1)
	ang := 0
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &ang)
		angryMax[i+1] = ang
	}
	l, r := 0, 0
	for i := 0; i < m; i++ {
		_, _ = fmt.Scanf("%d %d", &l, &r)
		if l == r {
			if angryMax[l] < i+1 {
				fmt.Printf("%d", i)
				return
			}
		}
		//for j := l; j <= r; j++ {
		q := angryMax[l : r+1]
		if findMinByDiv(&q) < i+1 {
			fmt.Printf("%d", i)
			return
		}
		//}
	}

	fmt.Printf("%d", m)
}
