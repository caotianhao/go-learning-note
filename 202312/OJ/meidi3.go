package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 1|家用电器,2|空调,3|冰箱,4|壁挂式空调,5|双开门冰箱
// 2#家用电器|冰箱|双开门冰箱,3#家用电器|空调|壁挂式空调
func totalSend(firstLine, secondLine string) {
	kind := make(map[int]string)
	sent := make(map[string]int)
	kindSlice := make([]int, 0)

	line1 := strings.Split(firstLine, ",")
	for _, v := range line1 {
		v0 := strings.Split(v, "|")
		id, _ := strconv.Atoi(v0[0])
		kindSlice = append(kindSlice, id)
		name := v0[1]
		kind[id] = name
	}

	line2 := strings.Split(secondLine, ",")
	for _, v := range line2 {
		id003 := strings.Index(v, "#")
		num, _ := strconv.Atoi(v[:id003])
		goods := strings.Split(v[id003+1:], "|")
		for _, g := range goods {
			sent[g] += num
		}
	}

	for _, v := range kindSlice {
		fmt.Printf("%d %d\n", v, sent[kind[v]])
	}
}

func main() {
	var firstLine, secondLine string
	_, _ = fmt.Scanf("%s", &firstLine)
	_, _ = fmt.Scanf("%s", &secondLine)
	totalSend(firstLine, secondLine)
}
