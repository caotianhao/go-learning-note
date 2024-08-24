package main

import "fmt"

type Point struct {
	X int
	Y int
}

func maxIncomeProducts(products []*Point, months int) (r []*Point) {
	// 什么叫投入月份的整数倍？
	// 3,6,9 那个用例，投入 11 个月，直接买 9 不就行？
	// 直接买 9 是 3654 收益，3+6 是 2775 收益？
	weight, value := make([]int, 0), make([]int, 0)
	finance := make(map[int]int)
	for _, product := range products {
		weight = append(weight, product.X)
		value = append(value, product.Y)
		finance[product.X] = product.Y
	}
	dp := make([][][]int, 0)
	for i := months; i >= 1; i-- {
		dp = append(dp, combinationSum(weight, i))
	}
	maxMoney := -1
	t := make([]int, 0)
	for _, monthCombination := range dp {
		for _, comb := range monthCombination {
			money := 0
			for _, val := range comb {
				money += finance[val]
			}
			if money > maxMoney {
				maxMoney = money
				t = comb
			}
		}
	}
	for _, v := range t {
		p := &Point{
			X: v,
			Y: finance[v],
		}
		r = append(r, p)
	}
	return
}

func combinationSum(candidates []int, target int) (ans [][]int) {
	comb := make([]int, 0)
	var dfs func(int, int)
	dfs = func(target int, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		dfs(target, idx+1)
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return
}

func main() {
	k1 := &Point{3, 902}
	k2 := &Point{6, 1873}
	k3 := &Point{9, 3654}
	fmt.Println(maxIncomeProducts([]*Point{k1, k2, k3}, 11))
}
