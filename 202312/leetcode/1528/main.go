package main

import "fmt"

func restoreString(s string, indices []int) string {
	str, l := "", len(s)
	map1528 := make(map[int]byte)
	for i := 0; i < l; i++ {
		map1528[indices[i]] = s[i]
	}
	for i := 0; i < l; i++ {
		str += string(map1528[i])
	}
	return str
}

func main() {
	s := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}
	fmt.Println(restoreString(s, indices))
}
