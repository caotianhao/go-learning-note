package main

import "fmt"

const max2761 = 1e6 + 1

var (
	prime2761      []int
	isNotPrime2761 [max2761]bool
)

func init() {
	for i := 2; i < max2761; i++ {
		if !isNotPrime2761[i] {
			prime2761 = append(prime2761, i)
			for j := i * i; j < max2761; j += i {
				isNotPrime2761[j] = true
			}
		}
	}
}

func findPrimePairs(n int) (res [][]int) {
	for _, x := range prime2761 {
		y := n - x
		if x > y {
			break
		}
		if !isNotPrime2761[y] {
			res = append(res, []int{x, y})
		}
	}
	return
}

func main() {
	fmt.Println(findPrimePairs(10))
	fmt.Println(findPrimePairs(7))
	fmt.Println(findPrimePairs(9))
	fmt.Println(findPrimePairs(11))
}
