package main

import "fmt"

func destCity(paths [][]string) string {
	l := len(paths)
	begin, end, cnt, des := make([]string, 0), make([]string, 0), 0, ""
	for i := 0; i < l; i++ {
		begin = append(begin, paths[i][0])
		end = append(end, paths[i][1])
	}
	for i := 0; i < l; i++ {
		cnt = 0
		for j := 0; j < l; j++ {
			if end[i] != begin[j] {
				cnt++
			}
			if cnt == l {
				des += end[i]
			}
		}
	}
	return des
}

func main() {
	path := [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}
	fmt.Println(destCity(path))
}
