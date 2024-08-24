package main

import "fmt"

func exist(board []string, word string) bool {
	row := len(board)
	col := len(board[0])
	boardInByte := make([][]byte, row)
	for i := range boardInByte {
		boardInByte[i] = make([]byte, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			boardInByte[i][j] = board[i][j]
		}
	}
	return helpExist(boardInByte, word)
}

func helpExist(boardInByte [][]byte, target string) bool {
	var result bool
	row, col := len(boardInByte), len(boardInByte[0])
	var dfs func(int, int, int)
	dfs = func(ii int, jj int, nxt int) {
		if ii < 0 || ii >= row || jj < 0 || jj >= col {
			return
		}
		if boardInByte[ii][jj] == '0' || boardInByte[ii][jj] != target[nxt] {
			return
		}
		if nxt == len(target)-1 {
			result = true
			return
		}
		tempBoard := boardInByte[ii][jj]
		boardInByte[ii][jj] = '0'
		dfs(ii-1, jj, nxt+1)
		dfs(ii+1, jj, nxt+1)
		dfs(ii, jj-1, nxt+1)
		dfs(ii, jj+1, nxt+1)
		boardInByte[ii][jj] = tempBoard
	}
	for ii := 0; ii < row; ii++ {
		for jj := 0; jj < col; jj++ {
			nxt := 0
			dfs(ii, jj, nxt)
		}
	}
	return result
}

func main() {
	fmt.Println(exist([]string{"XYZE", "SFZS", "XDEE"}, "XYZZED"))
}
