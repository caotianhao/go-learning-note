package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		isOne := make([]int, 9)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				isOne[int(board[i][j]-49)] += 1
			}
		}
		for i := 0; i < 9; i++ {
			if isOne[i] != 0 && isOne[i] != 1 {
				return false
			}
		}
	}

	for i := 0; i < 9; i++ {
		isOne := make([]int, 9)
		for j := 0; j < 9; j++ {
			if board[j][i] != '.' {
				isOne[int(board[j][i]-49)] += 1
			}
		}
		for i := 0; i < 9; i++ {
			if isOne[i] != 0 && isOne[i] != 1 {
				return false
			}
		}
	}

	isOne := make([]int, 9)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] != '.' {
				isOne[int(board[i][j]-49)] += 1
			}
		}
	}
	for i := 0; i < 9; i++ {
		if isOne[i] != 0 && isOne[i] != 1 {
			return false
		}
	}

	isOne = make([]int, 9)
	for i := 0; i < 3; i++ {
		for j := 3; j < 6; j++ {
			if board[i][j] != '.' {
				isOne[int(board[i][j]-49)] += 1
			}
		}
	}
	for i := 0; i < 9; i++ {
		if isOne[i] != 0 && isOne[i] != 1 {
			return false
		}
	}

	isOne = make([]int, 9)
	for i := 0; i < 3; i++ {
		for j := 6; j < 9; j++ {
			if board[i][j] != '.' {
				isOne[int(board[i][j]-49)] += 1
			}
		}
	}
	for i := 0; i < 9; i++ {
		if isOne[i] != 0 && isOne[i] != 1 {
			return false
		}
	}

	isOne = make([]int, 9)
	for i := 3; i < 6; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] != '.' {
				isOne[int(board[i][j]-49)] += 1
			}
		}
	}
	for i := 0; i < 9; i++ {
		if isOne[i] != 0 && isOne[i] != 1 {
			return false
		}
	}

	isOne = make([]int, 9)
	for i := 3; i < 6; i++ {
		for j := 3; j < 6; j++ {
			if board[i][j] != '.' {
				isOne[int(board[i][j]-49)] += 1
			}
		}
	}
	for i := 0; i < 9; i++ {
		if isOne[i] != 0 && isOne[i] != 1 {
			return false
		}
	}

	isOne = make([]int, 9)
	for i := 3; i < 6; i++ {
		for j := 6; j < 9; j++ {
			if board[i][j] != '.' {
				isOne[int(board[i][j]-49)] += 1
			}
		}
	}
	for i := 0; i < 9; i++ {
		if isOne[i] != 0 && isOne[i] != 1 {
			return false
		}
	}

	isOne = make([]int, 9)
	for i := 6; i < 9; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] != '.' {
				isOne[int(board[i][j]-49)] += 1
			}
		}
	}
	for i := 0; i < 9; i++ {
		if isOne[i] != 0 && isOne[i] != 1 {
			return false
		}
	}

	isOne = make([]int, 9)
	for i := 6; i < 9; i++ {
		for j := 3; j < 6; j++ {
			if board[i][j] != '.' {
				isOne[int(board[i][j]-49)] += 1
			}
		}
	}
	for i := 0; i < 9; i++ {
		if isOne[i] != 0 && isOne[i] != 1 {
			return false
		}
	}

	isOne = make([]int, 9)
	for i := 6; i < 9; i++ {
		for j := 6; j < 9; j++ {
			if board[i][j] != '.' {
				isOne[int(board[i][j]-49)] += 1
			}
		}
	}
	for i := 0; i < 9; i++ {
		if isOne[i] != 0 && isOne[i] != 1 {
			return false
		}
	}

	return true
}

func main() {
	arr := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	fmt.Println(isValidSudoku(arr))
}
