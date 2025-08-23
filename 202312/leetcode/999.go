package main

import "fmt"

func numRookCaptures(board [][]byte) int {
	rookR, rookC, zu := 0, 0, 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 'R' {
				rookR, rookC = i, j
			}
		}
	}
	//up
	for i := rookR; i > 0; i-- {
		k := 1
		if i-k >= 0 {
			if board[i-k][rookC] == 'B' {
				break
			}
			if board[i-k][rookC] == 'p' {
				zu++
				break
			}
		}
		k++
	}
	//down
	for i := rookR; i < 8; i++ {
		k := 1
		if i+k < 8 {
			if board[i+k][rookC] == 'B' {
				break
			}
			if board[i+k][rookC] == 'p' {
				zu++
				break
			}
		}
		k++
	}
	//left
	for j := rookC; j > 0; j-- {
		k := 1
		if j-k >= 0 {
			if board[rookR][j-k] == 'B' {
				break
			}
			if board[rookR][j-k] == 'p' {
				zu++
				break
			}
		}
		k++
	}
	//right
	for j := rookC; j < 8; j++ {
		k := 1
		if j+k < 8 {
			if board[rookR][j+k] == 'B' {
				break
			}
			if board[rookR][j+k] == 'p' {
				zu++
				break
			}
		}
		k++
	}
	return zu
}

func main() {
	fmt.Println(numRookCaptures([][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', 'R', '.', '.', '.', 'p'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
	}))
	fmt.Println(numRookCaptures([][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'p', 'p', '.', 'R', '.', 'p', 'B', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'B', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
	}))
}
