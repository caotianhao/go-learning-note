package main

import "fmt"

func smallestRepunitDivByK(k int) int {
	if k == 1 {
		return 1
	}
	if k%2 == 0 || k%5 == 0 {
		return -1
	}
	tmp, cnt := 1%k, 1
	for {
		tmp = (10*tmp + 1) % k
		cnt++
		if tmp == 0 {
			return cnt
		}
	}
}

func main() {
	fmt.Println(smallestRepunitDivByK(1111111111111111))
}
