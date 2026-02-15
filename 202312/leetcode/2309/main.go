package main

import (
	"fmt"
	"sort"
)

func greatestLetter(s string) string {
	map2309, slice2309 := map[byte]int{}, make([]byte, 0)
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			if _, ok := map2309[byte(v)]; ok {
				slice2309 = append(slice2309, byte(v))
				continue
			}
			map2309[byte(v)+32]++
		} else {
			if _, ok := map2309[byte(v)]; ok {
				slice2309 = append(slice2309, byte(v)-32)
				continue
			}
			map2309[byte(v)-32]++
		}
	}
	sort.Slice(slice2309, func(i, j int) bool {
		return slice2309[i] > slice2309[j]
	})
	if len(slice2309) == 0 {
		return ""
	}
	return string(slice2309[0])
}

func main() {
	fmt.Println(greatestLetter("lEeTcOdE"))
	fmt.Println(greatestLetter("arRAzFif"))
	fmt.Println(greatestLetter("AbCdEfGhIjK"))
}
