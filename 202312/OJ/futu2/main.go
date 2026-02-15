package main

import (
	"fmt"
	"sort"
)

func fightGetCoin(originAtk int, monster *[]int) int {
	coin := 0
	maxOfCoin := -1
	sort.Ints(*monster)
	if originAtk < (*monster)[0] {
		return 0
	}
	for i := 0; i < len(*monster); i++ {
		// 如果打的过，金币加一
		if originAtk >= (*monster)[i] {
			coin++
			maxOfCoin = maxFT(maxOfCoin, coin)
		} else {
			// 如果打不过
			// 打不过，且身上的钱加能力也打不过直接跑
			if originAtk+coin < (*monster)[i] {
				maxOfCoin = maxFT(maxOfCoin, coin)
				return maxOfCoin
			} else {
				// 金币换能力，扣除金币
				coin -= (*monster)[i] - originAtk
				// 提升能力
				originAtk += (*monster)[i] - originAtk
				// 打败怪物
				coin++
				maxOfCoin = maxFT(maxOfCoin, coin)
			}
		}
	}
	return maxOfCoin
}

func maxFT(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	originAtk, numberOfMonster := 0, 0
	_, _ = fmt.Scanf("%d %d", &originAtk, &numberOfMonster)
	atkOfMonster := 0
	monster := make([]int, 0)
	for i := 0; i < numberOfMonster; i++ {
		_, _ = fmt.Scanf("%d", &atkOfMonster)
		monster = append(monster, atkOfMonster)
	}
	fmt.Printf("%d", fightGetCoin(originAtk, &monster))
}
