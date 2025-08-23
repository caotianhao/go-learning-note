package main

import "fmt"

// 给你四个整数 length ，width ，height 和 mass ，分别表示一个箱子的三个维度和质量。
// 如果满足以下条件，那么箱子是 "Bulky" 的：
//
//	箱子 至少有一个 维度大于等于 1e4 。
//	或者箱子的 体积 大于等于 1e9 。
//
// 如果箱子的质量大于等于 100 ，那么箱子是 "Heavy" 的。

// 如果箱子同时是 "Bulky" 和 "Heavy" ，那么返回类别为 "Both" 。
// 如果箱子既不是 "Bulky" ，也不是 "Heavy" ，那么返回类别为 "Neither" 。
// 如果箱子是 "Bulky" 但不是 "Heavy" ，那么返回类别为 "Bulky" 。
// 如果箱子是 "Heavy" 但不是 "Bulky" ，那么返回类别为 "Heavy" 。

// 注意，箱子的体积等于箱子的长度、宽度和高度的乘积。
func categorizeBox(length, width, height, mass int) string {
	var isBulky, isHeavy bool
	v := int64(length) * int64(width) * int64(height)
	if length >= 10000 || width >= 10000 ||
		height >= 10000 || v >= 1000000000 {
		isBulky = true
	}
	if mass >= 100 {
		isHeavy = true
	}
	if isBulky && isHeavy {
		return "Both"
	} else if isBulky && !isHeavy {
		return "Bulky"
	} else if !isBulky && isHeavy {
		return "Heavy"
	} else {
		return "Neither"
	}
}

func main() {
	fmt.Println(categorizeBox(1000, 35, 700, 300))
}
