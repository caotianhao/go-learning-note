package main

import "fmt"

func memLeak(memory1, memory2 int) []int {
	i := 1
	for {
		if memory1 >= memory2 {
			memory1 -= i
			if memory1 < 0 {
				memory1 += i
				break
			}
		} else {
			memory2 -= i
			if memory2 < 0 {
				memory2 += i
				break
			}
		}
		i++
	}
	return []int{i, memory1, memory2}
}

func main() {
	fmt.Println(memLeak(2, 2))
	fmt.Println(memLeak(8, 11))
}
