package main

import (
	"fmt"
	"sort"
)

func filterRestaurants(restaurants [][]int, veganFriendly, maxPrice, maxDistance int) (ans []int) {
	sort.Slice(restaurants, func(i, j int) bool {
		return restaurants[i][1] > restaurants[j][1] ||
			restaurants[i][1] == restaurants[j][1] &&
				restaurants[i][0] > restaurants[j][0]
	})
	for _, r := range restaurants {
		if (r[3] <= maxPrice && r[4] <= maxDistance) &&
			(veganFriendly == 0 ||
				veganFriendly == 1 && r[2] == 1) {
			ans = append(ans, r[0])
		}
	}
	return
}

func main() {
	fmt.Println(filterRestaurants([][]int{
		{1, 4, 1, 40, 10},
		{2, 8, 0, 50, 5},
		{3, 8, 1, 30, 4},
		{4, 10, 0, 10, 3},
		{5, 1, 1, 15, 1},
	}, 0, 50, 10))
}
