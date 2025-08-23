package main

import (
	"fmt"
	"strconv"
)

func similarRGB(color string) string {
	hexArr := [16]int{0x00, 0x11, 0x22, 0x33, 0x44, 0x55,
		0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff}
	needArr := [16]string{"00", "11", "22", "33", "44", "55",
		"66", "77", "88", "99", "aa", "bb", "cc", "dd", "ee", "ff"}
	aa, _ := strconv.ParseInt(color[1:3], 16, 64)
	bb, _ := strconv.ParseInt(color[3:5], 16, 64)
	cc, _ := strconv.ParseInt(color[5:], 16, 64)
	minCC := 250000
	a, b, c := int(aa), int(bb), int(cc)
	ii, jj, kk := -1, -1, -1
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			for k := 0; k < 16; k++ {
				t := (a-hexArr[i])*(a-hexArr[i]) + (b-hexArr[j])*
					(b-hexArr[j]) + (c-hexArr[k])*(c-hexArr[k])
				if t < minCC {
					minCC = t
					ii, jj, kk = i, j, k
				}
			}
		}
	}
	return "#" + needArr[ii] + needArr[jj] + needArr[kk]
}

func main() {
	fmt.Println(similarRGB("#09f166"))
	fmt.Println(similarRGB("#4e3fe1"))
}
