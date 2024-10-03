package main

import (
	"fmt"
)

func main() {
	// 缺数据，做不了，回归分析也不会，缺个 25 和 33 的数据
	moji := make([]int, 36)
	moji[10] = 1705
	moji[11] = 1925
	moji[12] = 2200
	moji[13] = 2530
	moji[14] = 2860
	moji[15] = 3245
	moji[16] = 3630
	moji[17] = 4070
	moji[18] = 4510
	moji[19] = 5005
	moji[20] = 5500
	moji[21] = 5885
	moji[22] = 6875
	moji[23] = 7975
	moji[27] = 13970
	moji[30] = 20570
	moji[31] = 23155
	moji[35] = 34760

	for i := 0; i < 35; i++ {
		if moji[i] != 0 && moji[i+1] != 0 {
			fmt.Printf("%d - %d = %d\n", i+1, i, moji[i+1]-moji[i])
		}
	}

	//fmt.Println(23155 + 2585*4)
}
