package main

import "fmt"

func supplyWagon(supplies []int) []int {
	br := len(supplies) - len(supplies)/2
	for recycle := 0; recycle < br; recycle++ {
		minW := 3000
		for i := 0; i < len(supplies)-1; i++ {
			if supplies[i]+supplies[i+1] < minW {
				minW = supplies[i] + supplies[i+1]
			}
		}
		ind := -1
		for i := 0; i < len(supplies)-1; i++ {
			if supplies[i]+supplies[i+1] == minW {
				ind = i
				break
			}
		}
		ans := make([]int, 0)
		for i := 0; i < len(supplies); i++ {
			if i == ind {
				ans = append(ans, supplies[i]+supplies[i+1])
				continue
			}
			if i == ind+1 {
				continue
			}
			ans = append(ans, supplies[i])
		}
		supplies = ans
	}
	return supplies
}

func main() {
	fmt.Println(supplyWagon([]int{7, 3, 6, 1, 8}))
	fmt.Println(supplyWagon([]int{1, 3, 1, 5}))
}
