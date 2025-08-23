package main

import "fmt"

func plusOne(digits []int) []int {
	l := len(digits)
	slice66 := make([]int, 0)
	//最后一位不是 9，直接 +1
	if digits[l-1] != 9 {
		digits[l-1] += 1
		return digits
	} else {
		//是否全 9
		flag := true
		for i := 0; i < l; i++ {
			if digits[i] != 9 {
				flag = false
			}
		}
		//全 9
		if flag {
			slice66 = append(slice66, 1)
			for i := 0; i < l; i++ {
				slice66 = append(slice66, 0)
			}
			return slice66
		} else {
			//不是全9
			for i := l - 1; i >= 0; i-- {
				if digits[i] != 9 {
					for j := 0; j < i; j++ {
						slice66 = append(slice66, digits[j])
					}
					slice66 = append(slice66, digits[i]+1)
					for j := i + 1; j < l; j++ {
						slice66 = append(slice66, 0)
					}
					break
				}
			}
		}
	}
	return slice66
}

func main() {
	fmt.Println(plusOne([]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3,
		2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}))
	fmt.Println(plusOne([]int{8, 9, 9, 9}))
	fmt.Println(plusOne([]int{2, 4, 9, 3, 9}))
}
