package main

import "fmt"

func calculateTax(brackets [][]int, income int) float64 {
	l, tax, id, need := len(brackets), 0.0, -1, income
	if l == 1 {
		return float64(income) * float64(brackets[0][1]) / 100.0
	}
	for i := 0; i < l; i++ {
		if brackets[i][0] >= income {
			id = i
			break
		}
	}
	for i := 0; i < id; i++ {
		if i != 0 {
			tax += float64(brackets[i][0]-brackets[i-1][0]) * float64(brackets[i][1]) / 100.0
			continue
		}
		tax += float64(brackets[0][0]) * float64(brackets[0][1]) / 100.0
	}
	if id == 0 {
		return float64(income) * float64(brackets[0][1]) / 100.0
	} else {
		need -= brackets[id-1][0]
	}
	tax += float64(need) * float64(brackets[id][1]) / 100.0
	return tax
}

func main() {
	fmt.Println(calculateTax([][]int{{3, 50}, {7, 10}, {12, 25}}, 10))
	fmt.Println(calculateTax([][]int{{2, 50}}, 0))
	fmt.Println(calculateTax([][]int{{1, 0}, {4, 25}, {5, 50}}, 2))
	fmt.Println(calculateTax([][]int{{1, 75}, {5, 54}}, 1))
}
