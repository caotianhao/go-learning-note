package main

import "fmt"

func totalFruit(fruits []int) int {
	if len(fruits) <= 1 {
		return len(fruits)
	}
	p, q, fruitMap, ret := 0, 0, map[int]int{}, 0
	for q < len(fruits) {
		fruitMap[fruits[q]]++
		if len(fruitMap) <= 2 {
			ret = max(ret, q-p+1)
		}
		if len(fruitMap) == 3 {
			tmp := -1
			for i := range fruitMap {
				if i != fruits[q] && i != fruits[q-1] {
					delete(fruitMap, i)
					tmp = i
					break
				}
			}
			for j := q - 1; j >= 0; j-- {
				if fruits[j] == tmp {
					p = j + 1
					break
				}
			}
			ret = max(ret, q-p+1)
		}
		q++
	}
	return ret
}

func main() {
	fmt.Println(totalFruit([]int{1, 2, 1}))                         //3
	fmt.Println(totalFruit([]int{1, 2, 3, 2, 2}))                   //4
	fmt.Println(totalFruit([]int{1, 2, 2, 1, 3, 2, 1, 2, 2, 2}))    //5
	fmt.Println(totalFruit([]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4})) //5
	fmt.Println(totalFruit([]int{1, 1}))                            //2
}
