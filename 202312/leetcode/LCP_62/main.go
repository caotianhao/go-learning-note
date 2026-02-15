package main

import (
	"fmt"
)

func transportationHub(path [][]int) int {
	//point := make([]int, 0)
	//for _, v := range path {
	//	point = append(point, v[0])
	//	point = append(point, v[1])
	//}
	//sort.Ints(point)
	//pMax := point[len(point)-1] + 1
	//graph := make([][]int, pMax)
	//for i := 0; i < pMax; i++ {
	//	graph[i] = make([]int, pMax)
	//}
	//for _, v := range path {
	//	graph[v[0]][v[1]] = 1
	//}
	//res := -1
	//for i := 0; i < pMax; i++ {
	//	if isAllWayToPoint(i, pMax, graph) && isAllWayPointTo(i, pMax, graph) {
	//		res = i
	//	}
	//}
	//return res

	// out 出去的，也就是起点
	// in  到达的，也就是终点
	// all 统计站点数量
	out, in, all := map[int]bool{}, map[int]int{}, map[int]struct{}{}
	for _, v := range path {
		// 该站点有到达其他站点，设为 true
		out[v[0]] = true
		// 有其他站点来的到达本站点，加1
		in[v[1]]++
		// 所有站点存入，计算站点个数
		all[v[0]] = struct{}{}
		all[v[1]] = struct{}{}
	}
	// 站点个数
	l := len(all)
	for i, v := range in {
		// 如果站点不作为起点，且，其他所有站点都有到达该站点
		if !out[i] && v == l-1 {
			return i
		}
	}
	return -1
}

//func isAllWayToPoint(point, pMax int, graph [][]int) bool {
//	for i := 0; i < pMax; i++ {
//		if i == point {
//			continue
//		}
//		if graph[i][point] != 1 {
//			return false
//		}
//	}
//	return true
//}

//func isAllWayPointTo(point, pMax int, graph [][]int) bool {
//	for i := 0; i < pMax; i++ {
//		if graph[point][i] != 0 {
//			return false
//		}
//	}
//	return true
//}

func main() {
	fmt.Println(transportationHub([][]int{{0, 1}, {0, 3}, {1, 3}, {2, 0}, {2, 3}}))         //3
	fmt.Println(transportationHub([][]int{{0, 3}, {1, 0}, {1, 3}, {2, 0}, {3, 0}, {3, 2}})) //-1
	fmt.Println(transportationHub([][]int{{2, 5}, {4, 3}, {2, 3}}))                         //-1
}
