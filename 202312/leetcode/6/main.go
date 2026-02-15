package main

import "fmt"

func convert(s string, numRows int) string {
	l := len(s)
	// 如果 numRows 为 1，则不需要排列
	// 如果 numRows 大于等于 l，则本身就被排成一列
	if numRows == 1 || numRows >= l {
		return s
	}
	// 模拟每行存储
	arr := make([][]string, numRows)
	for i := range arr {
		arr[i] = make([]string, 0)
	}
	// 对应的行数
	idx := make([]int, 0)
	for i := 0; i < numRows; i++ {
		idx = append(idx, i)
	}
	for i := 0; i < numRows-2; i++ {
		idx = append(idx, numRows-i-2)
	}
	// 模拟
	for i := 0; i < l; i++ {
		arr[idx[i%(2*(numRows-1))]] = append(arr[idx[i%(2*(numRows-1))]], string(s[i]))
	}
	// 从左到右从上到下依次加起来
	str := ""
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			str += arr[i][j]
		}
	}
	return str
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert("PAYPALISHIRING", 2))
	fmt.Println(convert("P", 1))
}
