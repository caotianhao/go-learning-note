package main

import (
	"fmt"
	"math"
)

func isBoomerang(p [][]int) bool {
	// 判断三个点，需要满足 各不相同 且 不在同一条直线上
	// 变换为 判断三个点能否形成三角形
	// 即 三角形面积不为 0
	s1 := math.Sqrt(float64((p[1][1]-p[0][1])*(p[1][1]-p[0][1]) + (p[1][0]-p[0][0])*(p[1][0]-p[0][0])))
	s2 := math.Sqrt(float64((p[2][1]-p[0][1])*(p[2][1]-p[0][1]) + (p[2][0]-p[0][0])*(p[2][0]-p[0][0])))
	s3 := math.Sqrt(float64((p[1][1]-p[2][1])*(p[1][1]-p[2][1]) + (p[1][0]-p[2][0])*(p[1][0]-p[2][0])))
	pa := (s1 + s2 + s3) / 2
	// 判断浮点数是否为 0 应该这样判断
	// 同理若要判断两个浮点数是否相等，例如 14.99999996 和 15
	// 要判断其差的绝对值是否小于一个很小的数
	return math.Abs(pa*(pa-s1)*(pa-s2)*(pa-s3)) > 1e-9
}

func main() {
	fmt.Println(isBoomerang([][]int{{1, 1}, {2, 2}, {3, 3}}))       // false
	fmt.Println(isBoomerang([][]int{{1, 1}, {2, 3}, {3, 2}}))       // true
	fmt.Println(isBoomerang([][]int{{1, 1}, {2, 3}, {1, 1}}))       // false
	fmt.Println(isBoomerang([][]int{{73, 31}, {73, 19}, {73, 45}})) // false
	fmt.Println(isBoomerang([][]int{{22, 33}, {37, 27}, {67, 15}})) // false
	fmt.Println(isBoomerang([][]int{{0, 0}, {0, 2}, {2, 1}}))       // true
}
