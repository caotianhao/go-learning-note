package main

import (
	"fmt"
	"math"
)

//输入：time = [2,3,2], fruits = [[0,2],[1,4],[2,1]], limit = 3
//输出：10
//解释：
//由于单次最多采集 3 颗
//第 0 批需要采集 2 颗第 0 类型果实，需要采集 1 次，耗时为 2*1=2
//第 1 批需要采集 4 颗第 1 类型果实，需要采集 2 次，耗时为 3*2=6
//第 2 批需要采集 1 颗第 2 类型果实，需要采集 1 次，耗时为 2*1=2
//返回总耗时 2+6+2=10
func getMinimumTime(time []int, fruits [][]int, limit int) int {
	allTime, lf := 0, len(fruits)
	for i := 0; i < lf; i++ {
		allTime += time[fruits[i][0]] * int(math.Ceil(float64(fruits[i][1])/float64(limit)))
	}
	return allTime
}

func main() {
	//time := [2,3,2], fruits = [[0,2],[1,4],[2,1]], limit = 3
	fmt.Println(getMinimumTime([]int{2, 3, 2}, [][]int{{0, 2}, {1, 4}, {2, 1}}, 3))
}
