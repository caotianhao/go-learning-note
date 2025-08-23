package main

import "fmt"

func countLargestGroup(n int) (cnt int) {
	map1399, maxN := map[int][]int{}, -1
	for i := 1; i <= n; i++ {
		map1399[anySum1399(i)] = append(map1399[anySum1399(i)], i)
	}
	for _, v := range map1399 {
		if len(v) > maxN {
			maxN = len(v)
		}
	}
	for _, v := range map1399 {
		if len(v) == maxN {
			cnt++
		}
	}
	return
}

func anySum1399(n int) (sum int) {
	for n != 0 {
		sum += n % 10
		n /= 10
	}
	return
}

func main() {
	fmt.Println(countLargestGroup(39))
}
