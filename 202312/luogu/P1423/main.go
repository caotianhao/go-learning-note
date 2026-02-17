package main

import "fmt"

// 230 的由来
//
//	227                    228
//	 Σ   2*(0.98)^n < 99 ,  Σ   2*(0.98)^n > 99
//	n=0                    n=0
//
// 取近似 230
const s1423, p1423 = 0.98, 230

func main() {
	var m float64
	_, _ = fmt.Scan(&m)
	if m == 0.0 {
		fmt.Println(0)
		return
	}
	swim := make([]float64, 0)
	swim = append(swim, 2)
	for i := 0; i < p1423; i++ {
		swim = append(swim, swim[len(swim)-1]*s1423)
	}
	dis := float64(0)
	for i := range swim {
		dis += swim[i]
		if dis >= m {
			fmt.Println(i + 1)
			return
		}
	}
}
