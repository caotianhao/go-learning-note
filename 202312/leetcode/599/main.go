package main

import (
	"fmt"
	"sort"
)

func findRestaurant(list1 []string, list2 []string) (ret []string) {
	map1, map2 := map[string]int{}, map[string]int{}
	mapRest, mapRestV := map[string]int{}, make([]int, 0)
	for i, v := range list1 {
		map1[v] = i
	}
	for i, v := range list2 {
		map2[v] = i
	}
	for i, v := range map1 {
		if v2, ok := map2[i]; ok {
			mapRest[i] = v + v2
			mapRestV = append(mapRestV, v+v2)
		}
	}
	sort.Ints(mapRestV)
	for i, v := range mapRest {
		if v == mapRestV[0] {
			ret = append(ret, i)
		}
	}
	return
}

func main() {
	fmt.Println(findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		[]string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}))
}
