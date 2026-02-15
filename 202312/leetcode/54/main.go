package main

import "fmt"

func spiralOrder54(matrix [][]int) (ret []int) {
	row := len(matrix)
	if row == 0 {
		return
	}
	column := len(matrix[0])
	up, down, left, right, cnt := 0, row-1, 0, column-1, 0
loop:
	//螺旋矩阵没有算法
	//非常考验代码能力，就这个套路就行
	for cnt < row*column {
		for i := left; i <= right; i++ {
			ret = append(ret, matrix[up][i])
			cnt++
			//随时判断是否超过循环条件，超过的及时退出
			if cnt >= row*column {
				break loop
			}
		}
		up++
		for i := up; i <= down; i++ {
			ret = append(ret, matrix[i][right])
			cnt++
			if cnt >= row*column {
				break loop
			}
		}
		right--
		for i := right; i >= left; i-- {
			ret = append(ret, matrix[down][i])
			cnt++
			if cnt >= row*column {
				break loop
			}
		}
		down--
		for i := down; i >= up; i-- {
			ret = append(ret, matrix[i][left])
			cnt++
			if cnt >= row*column {
				break loop
			}
		}
		left++
	}
	return
}

func main() {
	fmt.Println(spiralOrder54([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
	fmt.Println(spiralOrder54([][]int{{1}}))
	fmt.Println(spiralOrder54([][]int{}))
}
