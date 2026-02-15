package main

import "fmt"

const maxNum = 60

var dpKS [maxNum][maxNum][maxNum][maxNum]int

func maxLengthWithCombineAAndB(a, b string) int {
	la, lb := len(a), len(b)
	ans := 1
	for p := 0; p <= la; p++ {
		for q := 0; q <= lb; q++ {
			for i := 0; i+p-1 < la; i++ {
				for k := 0; k+q-1 < lb; k++ {
					j := i + p - 1
					m := k + q - 1
					if p+q <= 1 {
						dpKS[i+1][j+1][k+1][m+1] = 1
					} else {
						dpKS[i+1][j+1][k+1][m+1] = 0
						if p > 1 && a[i] == a[j] {
							dpKS[i+1][j+1][k+1][m+1] |= dpKS[i+2][j][k+1][m+1]
						}
						if q > 1 && b[k] == b[m] {
							dpKS[i+1][j+1][k+1][m+1] |= dpKS[i+1][j+1][k+2][m]
						}
						if p > 0 && q > 0 && a[i] == b[m] {
							dpKS[i+1][j+1][k+1][m+1] |= dpKS[i+2][j+1][k+1][m]
						}
						if p > 0 && q > 0 && a[j] == b[k] {
							dpKS[i+1][j+1][k+1][m+1] |= dpKS[i+1][j][k+2][m+1]
						}
					}
					if dpKS[i+1][j+1][k+1][m+1] == 1 {
						ans = max(ans, p+q)
					}
				}
			}
		}
	}
	return ans
}

func main() {
	t := 0
	_, _ = fmt.Scanf("%d", &t)
	var a, b string
	for i := 0; i < t; i++ {
		_, _ = fmt.Scanf("%s", &a)
		_, _ = fmt.Scanf("%s", &b)
		fmt.Println(maxLengthWithCombineAAndB(a, b))
	}
}
