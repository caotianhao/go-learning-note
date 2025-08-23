package main

import "fmt"

func reformat(s string) string {
	letter, num := 0, 0
	letterArr, numArr := make([]string, 0), make([]string, 0)
	ans := ""
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			letter++
			letterArr = append(letterArr, string(v))
		} else if v >= '0' && v <= '9' {
			num++
			numArr = append(numArr, string(v))
		}
	}
	if abs1417(len(numArr), len(letterArr)) <= 1 {
		if len(numArr) > len(letterArr) {
			for i := 0; i < len(letterArr); i++ {
				ans += numArr[i]
				ans += letterArr[i]
			}
			ans += numArr[len(numArr)-1]
		} else if len(numArr) < len(letterArr) {
			for i := 0; i < len(numArr); i++ {
				ans += letterArr[i]
				ans += numArr[i]
			}
			ans += letterArr[len(letterArr)-1]
		} else {
			for i := 0; i < len(numArr); i++ {
				ans += letterArr[i]
				ans += numArr[i]
			}
		}
	}
	return ans
}

func abs1417(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	fmt.Println(reformat("a0b1c2"))
}
