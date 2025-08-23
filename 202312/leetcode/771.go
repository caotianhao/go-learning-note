package main

import "fmt"

func numJewelsInStones(jewels string, stones string) int {
	cnt := 0
	for _, v := range jewels {
		for i := 0; i < len(stones); i++ {
			if int32(stones[i]) == v {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
}
