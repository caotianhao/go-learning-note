package main

import "fmt"

func countSonMatrix(matrix [][]byte) (ans int) {
	//	r, c := len(matrix), len(matrix[0])
	//	for top := 0; top < r; top++ {
	//		for bottom := top; bottom < r; bottom++ {
	//			u := make(map[byte]bool)
	//			left := 0
	//			for right := 0; right < c; right++ {
	//				char := matrix[bottom][right]
	//				if _, ok := u[char]; ok {
	//					for matrix[top][left] != char {
	//						delete(u, matrix[top][left])
	//						left++
	//					}
	//					left++
	//				} else {
	//					u[char] = true
	//					ans += (right - left + 1) * (bottom - top + 1)
	//				}
	//			}
	//		}
	//	}
	//	return
	r, c := len(matrix), len(matrix[0])
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			for a := i; a < r; a++ {
				for b := j; b < c; b++ {
					t := countTotal(i, j, a, b, matrix)
					if t == 0 {
						break
					}
					ans += t
				}
			}
		}
	}
	return ans
}

func countTotal(a, b, c, d int, matrix [][]byte) int {
	hm := make(map[byte]int)
	for i := a; i <= c; i++ {
		for j := b; j <= d; j++ {
			hm[matrix[i][j]]++
			if hm[matrix[i][j]] > 1 {
				return 0
			}
		}
	}
	return 1
}

func main() {
	n, m := 0, 0
	_, _ = fmt.Scanf("%d %d", &n, &m)
	matrix := make([][]byte, n)
	for i := range matrix {
		matrix[i] = make([]byte, m)
	}
	var ys string
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%s", &ys)
		for idx, v := range ys {
			matrix[i][idx] = byte(v)
		}
	}
	fmt.Println(countSonMatrix(matrix))
}
