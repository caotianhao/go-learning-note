package main

import (
	"fmt"
)

func main() {
	// 缺数据，做不了，回归分析也不会，缺个 25 和 33 的数据
	baseAttack := make([]int, 36)
	baseAttack[10] = 3850
	baseAttack[11] = 4510
	baseAttack[12] = 5280
	baseAttack[13] = 6160
	baseAttack[14] = 7260
	baseAttack[15] = 8580
	baseAttack[16] = 10120
	baseAttack[17] = 11935
	baseAttack[18] = 14025
	baseAttack[19] = 16445
	baseAttack[20] = 19250
	baseAttack[21] = 22495
	baseAttack[22] = 26180
	baseAttack[23] = 30360
	baseAttack[27] = 52580
	baseAttack[30] = 75020
	baseAttack[31] = 83600
	baseAttack[35] = 123420

	for i := 0; i < 35; i++ {
		if baseAttack[i] != 0 && baseAttack[i+1] != 0 {
			fmt.Printf("%d - %d = %d\n", i+1, i, baseAttack[i+1]-baseAttack[i])
		}
	}
}
