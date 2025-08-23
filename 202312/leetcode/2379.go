package main

import "fmt"

func minimumRecolors(blocks string, k int) int {
	l, minN := len(blocks), k
	for i := 0; i < l-k; i++ {
		temp := howManyW2379(blocks[i : i+k])
		if temp < minN {
			minN = temp
		}
		if minN == 0 {
			return 0
		}
	}
	temp := howManyW2379(blocks[l-k:])
	if temp < minN {
		minN = temp
	}
	return minN
}

func howManyW2379(s string) int {
	cnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'W' {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(minimumRecolors("BBBBBBBBBBBBBBBB", 4))
	fmt.Println(minimumRecolors("WWWWWWWWWWWWWWWWWWWWWWWW", 4))
	fmt.Println(minimumRecolors("WBBWWBBWBW", 7))
	fmt.Println(minimumRecolors("WBWBBBW", 2))
}
