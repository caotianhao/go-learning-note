package main

import (
	"fmt"
)

var (
	dx = []int{-1, 1, 0, 0}
	dy = []int{0, 0, -1, 1}
)

func cc15(matrix [][]byte) int {
	rows := len(matrix)
	cols := len(matrix[0])
	if rows == 0 {
		return 0
	}
	queue := make([][2]int, 0)
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 's' {
				queue = append(queue, [2]int{i, j})
				visited[i][j] = true
			}
		}
	}
	steps := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
			x, y := curr[0], curr[1]
			if matrix[x][y] == 'e' {
				return steps
			}
			// 没截全的部分
			// 第一行判断边界
			// 第二行不为障碍物
			// 第三行没来过
			for k := 0; k < 4; k++ {
				newX, newY := x+dx[k], y+dy[k]
				if newX >= 0 && newX < rows && newY >= 0 && newY < cols &&
					matrix[newX][newY] != '#' &&
					!visited[newX][newY] {
					queue = append(queue, [2]int{newX, newY})
					visited[newX][newY] = true
				}
			}
		}
		steps++
	}
	return -1
}

func main() {
	fmt.Println(cc15([][]byte{
		{'.', 's', '.'},
		{'#', '#', '.'},
		{'e', '.', '.'},
	}))
}
