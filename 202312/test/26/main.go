package main

import (
	"fmt"
	"math"
	"sort"
)

func isIn(a, b, x, y int) bool {
	return a >= 0 && a < x && b >= 0 && b < y
}

func abs1(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func main() {
	notIn, outOfMax := 0, 0

	r, c := 0, 0
	_, _ = fmt.Scanf("%d %d", &r, &c)

	m := 0
	_, _ = fmt.Scanf("%d", &m)

	c1, c2 := 0, 0
	client := make([][]int, 0)
	for i := 0; i < m; i++ {
		_, _ = fmt.Scanf("%d %d", &c1, &c2)
		client = append(client, []int{c1, c2})
	}

	n := 0
	_, _ = fmt.Scanf("%d", &n)

	s1, s2 := 0, 0
	server := make([][]int, 0)
	for i := 0; i < m; i++ {
		_, _ = fmt.Scanf("%d %d", &s1, &s2)
		server = append(server, []int{s1, s2})
		if !isIn(s1, s2, r, c) {
			notIn++
		}
	}

	dis := make([]int, 0)
	for i := 0; i < len(server); i++ {
		t := math.MaxInt32
		for j := 0; j < len(client); j++ {
			sq := abs1(server[i][0]-client[j][0]) + abs1(server[i][1]-client[j][1])
			if sq > math.MaxInt32 {
				outOfMax++
			}
			if sq < t {
				t = sq
			}
		}
		dis = append(dis, t)
	}
	if notIn == n || outOfMax == n {
		fmt.Println(-1)
		return
	}
	sort.Ints(dis)
	fmt.Println(dis[0])
}
