package main

import "fmt"

func imageSmoother(img [][]int) [][]int {
	l, l2, point, sum := len(img), len(img[0]), 1, 0
	//go 创建二维数组必须以 int 型常数指定大小
	//但学到了！！可以以切片的形式这样创建
	//temp := make([][]int, l)
	//temp[i] = make([]int, l2)
	temp := make([][]int, l)
	for i := 0; i < l; i++ {
		temp[i] = make([]int, l2)
		for j := 0; j < l2; j++ {
			point, sum = 1, 0
			sum += img[i][j]
			if i-1 >= 0 && j-1 >= 0 {
				sum += img[i-1][j-1]
				point++
			}
			if i-1 >= 0 {
				sum += img[i-1][j]
				point++
			}
			if i-1 >= 0 && j+1 < l2 {
				sum += img[i-1][j+1]
				point++
			}
			if j-1 >= 0 {
				sum += img[i][j-1]
				point++
			}
			if j+1 < l2 {
				sum += img[i][j+1]
				point++
			}
			if i+1 < l && j-1 >= 0 {
				sum += img[i+1][j-1]
				point++
			}
			if i+1 < l {
				sum += img[i+1][j]
				point++
			}
			if i+1 < l && j+1 < l2 {
				sum += img[i+1][j+1]
				point++
			}
			temp[i][j] = sum / point
		}
	}
	return temp
}

func main() {
	fmt.Println(imageSmoother([][]int{
		{100, 200, 100},
		{200, 50, 200},
		{100, 200, 100},
	}))
}
