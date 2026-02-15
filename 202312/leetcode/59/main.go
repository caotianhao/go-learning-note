package main

import "fmt"

func generateMatrix(n int) [][]int {
	//创建n*n空矩阵
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}
	//记录上下左右的边界
	//cnt是添加的元素，同时需要判断边界
	up, down, left, right, cnt := 0, n-1, 0, n-1, 1
	for cnt <= n*n {
		//添加第一行，从左至右，添加完行加一
		for i := left; i <= right; i++ {
			ret[up][i] = cnt
			cnt++
		}
		//up++，此时up, down, left, right = 1, n-1, 0, n-1
		//相当于下次如果还能到这层循环的话，要从第2层开始加
		//同理
		up++
		for i := up; i <= down; i++ {
			ret[i][right] = cnt
			cnt++
		}
		right--
		for i := right; i >= left; i-- {
			ret[down][i] = cnt
			cnt++
		}
		down--
		for i := down; i >= up; i-- {
			ret[i][left] = cnt
			cnt++
		}
		left++
	}
	return ret
}

func main() {
	a := generateMatrix(13)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			fmt.Printf("%d\t", a[i][j])
		}
		fmt.Println()
	}
}
