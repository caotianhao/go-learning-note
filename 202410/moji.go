package main

import (
	"fmt"
)

func main() {
	// 缺数据，做不了，回归分析也不会，缺个 25 和 33 的数据
	additionalAttack := make([]int, 36)
	additionalAttack[10] = 1705
	additionalAttack[11] = 1925
	additionalAttack[12] = 2200
	additionalAttack[13] = 2530
	additionalAttack[14] = 2860
	additionalAttack[15] = 3245
	additionalAttack[16] = 3630
	additionalAttack[17] = 4070
	additionalAttack[18] = 4510
	additionalAttack[19] = 5005
	additionalAttack[20] = 5500
	additionalAttack[21] = 5885
	additionalAttack[22] = 6875
	additionalAttack[23] = 7975
	additionalAttack[27] = 13970
	additionalAttack[30] = 20570
	additionalAttack[31] = 23155
	additionalAttack[35] = 34760

	for i := 0; i < 35; i++ {
		if additionalAttack[i] != 0 && additionalAttack[i+1] != 0 {
			fmt.Printf("%d - %d = %d\n", i+1, i, additionalAttack[i+1]-additionalAttack[i])
		}
	}

	//fmt.Println(23155 + 2585*4)
}
